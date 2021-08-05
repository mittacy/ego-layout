package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"time"
)

// NewLogger 创建日志新句柄
// @param name 日志标识，用来区分多个日志，传空则为default日志文件
// @return *zap.Logger
func NewLogger(name string) *zap.Logger {
	if name == "" {
		name = bizDefaultLogName
	}

	zapConf := ZapConf{
		ServerName:   conf.ServerName,
		Path:         conf.Path,
		Name:         name,
		JsonFormat:   true,
		RotationTime: 24,
		MaxAge:       conf.BizMaxAge,
		LowLevel:     parseLevel(conf.LowLevel),
		HighLevel:    zapcore.FatalLevel,
	}

	// 请求日志特殊处理
	if name == RequestLogName {
		zapConf.MaxAge = conf.CallMaxAge
	}

	// 绑定服务名和服务环境
	if err := viper.UnmarshalKey("server.name", &zapConf.ServerName); err != nil {
		panic(fmt.Sprintf("new logger err: %s", err))
	}

	zapConf.CheckConf()

	lowCore, err := newCore(zapConf)
	if err != nil {
		panic(fmt.Sprintf("logger err: %v", err))
	}

	coreArr := []zapcore.Core{lowCore}
	return zap.New(zapcore.NewTee(coreArr...), zap.AddStacktrace(zapcore.WarnLevel), zap.Fields(zapcore.Field{
		Key:    "module_name",
		Type:   zapcore.StringType,
		String: zapConf.ServerName,
	}))
}

// newCore zapCore配置
// @param conf 配置信息
// @return zapcore.Core zapCore
// @return error
func newCore(conf ZapConf) (zapcore.Core, error) {
	priority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= conf.LowLevel && lev <= conf.HighLevel
	})

	writer, err := getWriter(conf.Path, conf.Name, conf.RotationTime, conf.MaxAge)
	if err != nil {
		return nil, err
	}

	var syncer zapcore.WriteSyncer
	if conf.LogInConsole {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writer))
	} else {
		syncer = zapcore.AddSync(writer)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "log_at",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "context",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 大写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,  // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	var encoder zapcore.Encoder
	if conf.JsonFormat {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		syncer,
		priority,
	)

	return core, nil
}

// getWriter 获取日志分割 Writer
// @param name 日志目录+名字, eg: ./logs/info/default
// @param rotationTime 分割日志时间，单位：小时
// @param maxAge 日志保留天数
// @return zapcore.WriteSyncer
func getWriter(dir, name string, rotationTime, maxAge time.Duration) (io.Writer, error) {
	dir = strings.Trim(dir, "/")
	file := fmt.Sprintf("%s/%s", dir, name)

	if rotationTime < 24 {
		file += "_%Y-%m-%d-H.log"
	} else {
		file += "_%Y-%m-%d.log"
	}

	writer, err := rotatelogs.New(
		file,
		rotatelogs.WithLinkName(dir+"/."+name),
		rotatelogs.WithMaxAge(time.Hour*24*maxAge),
		rotatelogs.WithRotationTime(time.Hour*rotationTime),
	)

	if err != nil {
		return nil, err
	}

	return writer, nil
}

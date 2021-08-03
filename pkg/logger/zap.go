package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/mittacy/ego-layout/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

func NewLogger(name string) *zap.Logger {
	// 低级日志 Core
	zapConf := ZapConf{
		ServerName:   config.Global.Server.Name,
		Path:         conf.Path + "/info/",
		Name:         name,
		JsonFormat:   true,
		RotationTime: 24,
		MaxAge:       conf.BizMaxAge,
		MinLevel:     parseLevel(conf.MinLevel),
		HighLevel:    zapcore.WarnLevel,
	}
	zapConf.CheckConf()

	lowCore, err := newCore(zapConf)
	if err != nil {
		panic(fmt.Sprintf("logger err: %v", err))
	}

	// 高级日志 Core
	zapConf.Path = conf.Path + "/error/"
	zapConf.MinLevel = zapcore.ErrorLevel
	zapConf.MaxAge = conf.BizErrMaxAge
	zapConf.CheckConf()

	highCore, err := newCore(zapConf)
	if err != nil {
		panic(fmt.Sprintf("logger err: %v", err))
	}

	coreArr := []zapcore.Core{lowCore, highCore}

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
		return lev >= conf.MinLevel && lev <= conf.HighLevel
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
// @param path 日志目录地址
// @param name 日志标识名，用来区分多个日志
// @param rotationTime 分割日志时间，单位：小时
// @param maxAge 日志保留天数
// @return zapcore.WriteSyncer
func getWriter(path, name string, rotationTime, maxAge time.Duration) (io.Writer, error) {
	fmt.Println("path: ", path)
	if rotationTime < 24 {
		path += name + "_%Y-%m-%d-H.log"
	} else {
		path += name + "_%Y-%m-%d.log"
	}
	fmt.Println("path: ", path)

	writer, err := rotatelogs.New(
		path,
		rotatelogs.WithLinkName(name),
		rotatelogs.WithMaxAge(time.Hour*24*maxAge),
		rotatelogs.WithRotationTime(time.Hour*rotationTime),
	)

	if err != nil {
		return nil, err
	}

	return writer, nil
}

package logger

import (
	"fmt"
	"github.com/mittacy/ego-layout/pkg/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const (
	jsonMarshalErrTitle = "json marshal err"
	copierErrTitle      = "copier err"
	transformErrTitle   = "transform err"
	cacheErrTitle       = "cache err"
)

// NewLogger 创建日志新句柄
// @param name 日志标识，用来区分多个日志，传空则为default日志文件
// @return *zap.Logger
func NewLogger(name string) *zap.Logger {
	logInConsole := false
	if gin.Mode() == gin.DebugMode {
		logInConsole = true
	}

	// warn、err日志记录到err文件夹
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // error
		return lev >= zap.ErrorLevel
	})
	highLogger := getLjLogger(name, true, config.Global.Log)
	highCore := newCore(highLogger, true, logInConsole, highPriority)

	// info、debug记录到info文件夹
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // info 和 debug
		if gin.Mode() == gin.DebugMode {
			return lev < zap.ErrorLevel && lev >= zap.DebugLevel
		}
		return lev < zap.ErrorLevel && lev >= zap.InfoLevel
	})
	lowLogger := getLjLogger(name, false, config.Global.Log)
	lowCore := newCore(lowLogger, true, logInConsole, lowPriority)

	coreArr := []zapcore.Core{highCore, lowCore}

	return zap.New(zapcore.NewTee(coreArr...), zap.AddStacktrace(zapcore.WarnLevel), zap.Fields(zapcore.Field{
		Key:    "module_name",
		Type:   zapcore.StringType,
		String: config.Global.Server.Name,
	}))
}

// CopierErrLog 结构体转化错误
// @param err
func CopierErrLog(err error) {
	LogWithStack(copierErrTitle, err)
}

// TransformErrLog 响应包装错误
// @param err
func TransformErrLog(err error) {
	LogWithStack(transformErrTitle, err)
}

// JsonMarshalErrLog json序列化与反序列化日志错误
// @param err
func JsonMarshalErrLog(err error) {
	LogWithStack(jsonMarshalErrTitle, err)
}

// CacheErrLog 缓存错误日志
// @param err
func CacheErrLog(err error) {
	LogWithStack(cacheErrTitle, err)
}

// LogWithStack 写带有调用栈的日志
// @param title 标题
// @param err 错误信息
func LogWithStack(title string, err error) {
	zap.L().Error(fmt.Sprintf("%s:%s", title, err),
		zap.String("trace", fmt.Sprintf("%+v", err)))
}

// newCore 获取日志配置结构
// @param lumber 日志压缩配置
// @param jsonFormat 是否输出为json格式
// @param logInConsole 是否同时输出到控制台
// @param levelEnable 该日志文件显示的级别: debug/info/warn/error
// @return zapcore.Core zap配置
func newCore(lumber *lumberjack.Logger, jsonFormat, logInConsole bool, levelEnable zap.LevelEnablerFunc) zapcore.Core {
	var syncer zapcore.WriteSyncer
	if logInConsole {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumber))
	} else {
		syncer = zapcore.AddSync(lumber)
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
	if jsonFormat {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		syncer,
		levelEnable,
	)

	return core
}

// getLjLogger 获取lumberjack 分割配置
// @param name 日志标识，用来区分多个日志，传空则写入整个服务共用的日志文件
// @param isErrLog 是否为错误日志
// @param logConfig 日志分割配置
// @return zapcore.WriteSyncer
func getLjLogger(name string, isErrLog bool, logConfig config.Log) *lumberjack.Logger {
	filePath := ""

	if isErrLog {
		logConfig.Path += "/err"
	} else {
		logConfig.Path += "/info"
	}

	if name != "" {
		filePath = fmt.Sprintf("%s/%s.log", logConfig.Path, name)
	} else {
		filePath = fmt.Sprintf("%s/default.log", logConfig.Path)
	}

	return &lumberjack.Logger{
		Filename:   filePath,             // 日志输出文件
		MaxSize:    logConfig.MaxSize,    // 日志最大保存1M
		MaxBackups: logConfig.MaxBackups, // 就日志保留5个备份
		MaxAge:     logConfig.MaxAge,     // 最多保留30个日志 和MaxBackups参数配置1个就可以
		Compress:   logConfig.Compress,   // 自导打 gzip包 默认false
	}
}

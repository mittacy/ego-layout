package log

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestDefault(t *testing.T) {
	InitStd()

	// 日志全局配置
	globalFields := []zapcore.Field{
		{
			Key:    "module_name",
			Type:   zapcore.StringType,
			String: "serverName",
		},
	}

	SetGlobalConfig(WithPath("./"), WithLevel(zapcore.DebugLevel), WithLogInConsole(true),
		WithGlobalFields(globalFields...), WithGlobalEncoderJSON(true))

	// 打印日志
	Debug("this is Debug")
	Info("this is Info")
	Warn("this is Warn")
	Error("this is Error")

	std.Debug("this is SugarDebug")
	std.Info("this is SugarInfo")
	std.Warn("this is SugarWarn")
	std.Error("this is SugarError")

	std.Debugf("this is %s", "Debugf")
	std.Infof("this is %s", "Infof")
	std.Warnf("this is %s", "Warn")
	std.Errorf("this is %s", "Errorf")

	std.Debugw("this is Debugw", "k", "Debugw")
	std.Infow("this is Infow", "k", "Infow")
	std.Warnw("this is Warnw", "k", "Warnw")
	std.Errorw("this is Errorw", "k", "Errorw")
}

func TestNewLog(t *testing.T) {
	InitStd()

	// 日志全局配置
	globalFields := []zapcore.Field{
		{
			Key:    "module_name",
			Type:   zapcore.StringType,
			String: "serverName",
		},
	}

	SetGlobalConfig(WithPath("./"), WithLevel(zapcore.DebugLevel), WithLogInConsole(true), WithGlobalFields(globalFields...))

	// 打印日志
	bizLog := NewWithLevel("testNew", zapcore.InfoLevel, zap.AddStacktrace(zapcore.WarnLevel))

	bizLog.Debug("this is SugarDebug")
	bizLog.Info("this is SugarInfo")
	bizLog.Warn("this is SugarWarn")
	bizLog.Error("this is SugarError")

	bizLog.Debugf("this is %s", "Debugf")
	bizLog.Infof("this is %s", "Infof")
	bizLog.Warnf("this is %s", "Warn")
	bizLog.Errorf("this is %s", "Errorf")

	bizLog.Debugw("this is Debugw", "k", "Debugw")
	bizLog.Infow("this is Infow", "k", "Infow")
	bizLog.Warnw("this is Warnw", "k", "Warnw")
	bizLog.Errorw("this is Errorw", "k", "Errorw")
}

func TestBizLog(t *testing.T) {
	InitStd()

	// 日志全局配置
	globalFields := []zapcore.Field{
		{
			Key:    "module_name",
			Type:   zapcore.StringType,
			String: "serverName",
		},
	}

	SetGlobalConfig(WithPath("./"), WithLevel(zapcore.DebugLevel), WithLogInConsole(true),
		WithGlobalFields(globalFields...), WithGlobalEncoderJSON(true))

	// 打印日志
	bizLog := NewWithLevel("testBiz", zapcore.InfoLevel, zap.AddStacktrace(zapcore.WarnLevel))

	bizLog.CopierErrLog(errors.New("copier error message"), nil)
	bizLog.TransformErrLog(errors.New("transform error message"), nil)
	bizLog.JsonMarshalErrLog(errors.New("json marshal error message"), nil)
	bizLog.CacheErrLog(errors.New("cache error message"), nil)
	bizLog.MysqlErrLog(errors.New("mysql error message"), nil)
	bizLog.UnknownErrLog(errors.New("unknown error message"), nil)
}

func TestConsoleLog(t *testing.T) {
	InitStd()

	// 日志全局配置
	globalFields := []zapcore.Field{
		{
			Key:    "module_name",
			Type:   zapcore.StringType,
			String: "serverName",
		},
	}

	SetGlobalConfig(WithPath("./"), WithLevel(zapcore.DebugLevel), WithLogInConsole(true),
		WithGlobalFields(globalFields...), WithGlobalEncoderJSON(false))

	// 打印日志
	bizLog := NewWithLevel("testBiz", zapcore.InfoLevel, zap.AddStacktrace(zapcore.WarnLevel))

	bizLog.CopierErrLog(errors.New("copier error message"), nil)
	bizLog.TransformErrLog(errors.New("transform error message"), nil)
	bizLog.JsonMarshalErrLog(errors.New("json marshal error message"), nil)
	bizLog.CacheErrLog(errors.New("cache error message"), nil)
	bizLog.MysqlErrLog(errors.New("mysql error message"), nil)
}

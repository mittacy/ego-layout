package log

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestDefault(t *testing.T) {
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
	Debug("this is Debug")
	Info("this is Info")
	Warn("this is Warn")
	Error("this is Error")

	std.Sugar().Debug("this is SugarDebug")
	std.Sugar().Info("this is SugarInfo")
	std.Sugar().Warn("this is SugarWarn")
	std.Sugar().Error("this is SugarError")

	std.Sugar().Debugf("this is SugarDebugf %s", "Debugf")
	std.Sugar().Infof("this is SugarInfo %s", "Infof")
	std.Sugar().Warnf("this is SugarWarnf %s", "Warn")
	std.Sugar().Errorf("this is SugarErrorf %s", "Errorf")
}

func TestNewLog(t *testing.T) {
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

	bizLog.Debug("this is Debug")
	bizLog.Info("this is Info")
	bizLog.Warn("this is Warn")
	bizLog.Error("this is Error")

	bizLog.Sugar().Debug("this is SugarDebug")
	bizLog.Sugar().Info("this is SugarInfo")
	bizLog.Sugar().Warn("this is SugarWarn")
	bizLog.Sugar().Error("this is SugarError")

	bizLog.Sugar().Debugf("this is SugarDebugf %s", "Debugf")
	bizLog.Sugar().Infof("this is SugarInfo %s", "Infof")
	bizLog.Sugar().Warnf("this is SugarWarnf %s", "Warn")
	bizLog.Sugar().Errorf("this is SugarErrorf %s", "Errorf")
}

func TestBizLog(t *testing.T) {
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
	bizLog := NewWithLevel("testBiz", zapcore.InfoLevel, zap.AddStacktrace(zapcore.WarnLevel))

	bizLog.CopierErrLog(errors.New("copier error message"))
	bizLog.TransformErrLog(errors.New("transform error message"))
	bizLog.JsonMarshalErrLog(errors.New("json marshal error message"))
	bizLog.CacheErrLog(errors.New("cache error message"))
	bizLog.MysqlErrLog(errors.New("mysql error message"))
}

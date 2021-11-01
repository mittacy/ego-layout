package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestZap(t *testing.T) {
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
	testLog := std

	testLog.Debug("this is Debug")
	testLog.Info("this is Info")
	testLog.Warn("this is Warn")
	testLog.Error("this is Error")

	testLog.Sugar().Debug("this is SugarDebug")
	testLog.Sugar().Info("this is SugarInfo")
	testLog.Sugar().Warn("this is SugarWarn")
	testLog.Sugar().Error("this is SugarError")

	testLog.Sugar().Debugf("this is SugarDebugf %s", "Debugf")
	testLog.Sugar().Infof("this is SugarInfo %s", "Infof")
	testLog.Sugar().Warnf("this is SugarWarnf %s", "Warn")
	testLog.Sugar().Errorf("this is SugarErrorf %s", "Errorf")
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

package log

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	logPath        = "."
	globalLowLevel = zapcore.DebugLevel
	logInConsole   = false
	globalFields   = make([]zapcore.Field, 0)
	jsonEncoder    = true // 是否为json格式日志
)

type ConfigOption func()

// SetGlobalConfig 设置日志全局配置
// @param options
func SetGlobalConfig(options ...ConfigOption) {
	for _, option := range options {
		option()
	}
}

// SetPath 设置日志路径, 修改后，新建的日志将会是新配置，已经建立的日志配置不变
// @param path 路径
func WithPath(path string) ConfigOption {
	return func() {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.MkdirAll(path, 0700); err != nil {
				fmt.Printf("log: create %s directory err: %s\n", path, err)
			}
		}

		logPath = path
		ResetDefault(InitStd())
	}
}

// SetLowLevel 设置服务记录的最低日志级别
// 修改后，新建的日志将会是新配置，已经建立的日志配置不变
// @param l 日志级别(-1:debug、0:info、1:warn、2:error)
func WithLevel(l zapcore.Level) ConfigOption {
	return func() {
		globalLowLevel = l
		ResetDefault(InitStd())
	}
}

// SetLogInConsole 是否输出到控制台
// 修改后，新建的日志将会是新配置，已经建立的日志配置不变
// @param isLogInConsole
func WithLogInConsole(isLogInConsole bool) ConfigOption {
	return func() {
		logInConsole = isLogInConsole
		ResetDefault(InitStd())
	}
}

// AddGlobalFields 添加全局日志的新字段, 新建的日志将会是新配置，已经建立的日志配置不变
// @param field 日志字段
func WithGlobalFields(field ...zapcore.Field) ConfigOption {
	return func() {
		globalFields = append(globalFields, field...)
		ResetDefault(InitStd())
	}
}

// WithGlobalEncoderJSON 是否为json格式日志
// @param isJSONEncoder
// @return ConfigOption
func WithGlobalEncoderJSON(isJSONEncoder bool) ConfigOption {
	return func() {
		jsonEncoder = isJSONEncoder
		ResetDefault(InitStd())
	}
}

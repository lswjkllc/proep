package commons

import (
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	// us "github.com/lswjkllc/proep/src/utils"
)

// 定义 Logger
var Logger *zap.Logger

func init() {
	// 初始化 Logger
	initLogger()
}

func initLogger() {
	/// 获取配置
	config := GetConfig()
	// 检查日志等级
	level := checkLevel(config.CommonBase.LogLevel)
	// 获取序列化器
	encoder := getEncoder(config)
	// 获取输出器
	sync := getWriteSync(config)
	// 获取核心数据结构
	core := zapcore.NewCore(encoder, sync, level)
	// 初始化
	Logger = zap.New(core)
}

// 负责日志写入的位置
func getWriteSync(config *ConfigInfo) zapcore.WriteSyncer {
	// 打开文件
	logPath := filepath.Join(config.CommonBase.LogPath, "/server.log")
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	// 输出到 文件
	syncFile := zapcore.AddSync(logFile)
	// 输出到 标准错误
	syncConsole := zapcore.AddSync(os.Stderr)
	// 返回 syncer
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}

// 设置序列化方式 (日志格式)
func getEncoder(config *ConfigInfo) zapcore.Encoder {
	// 获取 debug 信息
	debug := config.CommonBase.Debug
	// 获取 序列化配置
	encoderConfig := getEncoderConfig(debug)
	// 初始化 Encoder
	encoder := chooseEncoder(debug, encoderConfig)
	// 返回
	return encoder
}

// 选择序列化器
func chooseEncoder(debug bool, encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
	// 定义
	var encoder zapcore.Encoder
	// 初始化
	if debug {
		// debug 模型下, 使用 ConsoleEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		// 否则, 使用 JsonEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}
	return encoder
}

// 序列化配置
func getEncoderConfig(debug bool) zapcore.EncoderConfig {
	// 定义
	var encoderConfig zapcore.EncoderConfig
	// 初始化
	if debug {
		// debug 模式下, 使用 开发 配置
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		// 否则, 使用 生产 配置
		encoderConfig = zap.NewProductionEncoderConfig()
	}
	return encoderConfig
}

// 检查日志等级
func checkLevel(levelstr string) zapcore.Level {
	// 定义
	var level zapcore.Level
	// 将字符串转换为大写
	upperLevelstr := strings.ToUpper(levelstr)
	switch upperLevelstr {
	case zapcore.DebugLevel.CapitalString():
		// DEBUG
		level = zapcore.DebugLevel
	case zapcore.InfoLevel.CapitalString():
		// INFO
		level = zapcore.InfoLevel
	case zapcore.WarnLevel.CapitalString():
		// WARN
		level = zapcore.WarnLevel
	case zapcore.ErrorLevel.CapitalString():
		// ERROR
		level = zapcore.ErrorLevel
	case zapcore.DPanicLevel.CapitalString():
		// DPANIC
		level = zapcore.DPanicLevel
	case zapcore.PanicLevel.CapitalString():
		// PANIC
		level = zapcore.PanicLevel
	case zapcore.FatalLevel.CapitalString():
		// FATAL
		level = zapcore.FatalLevel
	default:
		panic("invalid log level => " + levelstr)
	}
	return level
}

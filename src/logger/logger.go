package logger

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	coms "github.com/lswjkllc/proep/src/commons"
	us "github.com/lswjkllc/proep/src/utils"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 定义 Logger
var Logger *zap.Logger

func init() {
	// 初始化 Logger
	initLogger()
}

func initLogger() {
	/// 获取配置
	config := coms.GetConfig()
	// 检查日志等级
	level := checkLevel(config.LogBase.Level)
	// 获取序列化器
	encoder := getEncoder(config)
	// 获取输出器
	sync := getWriteSync(config.LogBase)
	// 获取核心数据结构
	core := zapcore.NewCore(encoder, sync, level)
	// 堆栈跟踪 (添加日志行号)
	// caller := zap.AddCaller()
	// 初始化
	Logger = zap.New(core) //.WithOptions(caller)
}

// 负责日志写入的位置
func getWriteSync(logConfig coms.LogBaseEntity) zapcore.WriteSyncer {
	// 初始化
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logConfig.Path,       // 日志文件路径, 默认 os.TempDir()
		MaxSize:    logConfig.MaxSize,    // 每个日志文件保存10M, 默认 100M
		MaxBackups: logConfig.MaxBackups, // 保留30个备份, 默认不限
		MaxAge:     logConfig.MaxAge,     // 保留7天, 默认不限
		Compress:   logConfig.Compress,   // 是否压缩, 默认不压缩
	}
	// 输出到 文件
	syncFile := zapcore.AddSync(lumberJackLogger)
	// 输出到 标准错误
	syncConsole := zapcore.AddSync(os.Stderr)
	// 返回 syncer
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}

// 设置序列化方式 (日志格式)
func getEncoder(config *coms.ConfigInfo) zapcore.Encoder {
	// 获取 debug 信息
	debug := config.CommonBase.Debug
	// 获取 序列化配置
	encoderConfig := getEncoderConfig(debug)
	// 初始化 Encoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	// 返回
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

func Error(c echo.Context, msg string, fields ...zap.Field) {
	// 添加额外信息
	fields = addFields(c, fields...)
	// 输出
	Logger.Error(msg, fields...)
}

func Info(c echo.Context, msg string, fields ...zap.Field) {
	// 添加额外信息
	fields = addFields(c, fields...)
	// 输出
	Logger.Info(msg, fields...)
}

func addFields(c echo.Context, fields ...zap.Field) []zap.Field {
	// 获取 caller 链路信息
	caller := getCaller(2)
	// 获取 链路Id
	traceId := c.Request().Header.Get("trace-id")
	// 保存
	fields = append(fields, zap.String("caller", caller), zap.String("traceId", traceId))

	return fields
}

func getCaller(skipOff int) string {
	// 获取 调用信息
	_, filePath, line, ok := runtime.Caller(skipOff + 1)
	if !ok {
		panic("runtime caller error ...")
	}
	// 将 int 转 string
	sline := strconv.Itoa(line)
	// 连接并返回
	return us.JoinStrings(filePath, ":", sline)
}

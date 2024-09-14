package bootstrap

import (
	"github.com/giles-wong/roadShow/global"
	"github.com/giles-wong/roadShow/utils/tools"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var (
	level  zapcore.Level
	option []zap.Option
)

func InitLog() *zap.Logger {
	// 创建根目录
	createRootDir()
	// 设置日志等级
	setLogLevel()

	if global.App.Config.LogConf.ShowLine {
		option = append(option, zap.AddCaller())
	}

	// 初始化 zap
	return zap.New(getZapCore(), option...)
}

// 判断文件夹是否存在
func createRootDir() {
	if ok, _ := tools.PathExists(global.App.Config.LogConf.RootDir); !ok {
		_ = os.MkdirAll(global.App.Config.LogConf.RootDir, os.ModePerm)
	}
}

func setLogLevel() {
	switch global.App.Config.LogConf.Level {
	case "debug":
		level = zap.DebugLevel
		option = append(option, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		option = append(option, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(global.App.Config.AppConf.Env + "." + l.String())
	}

	// 设置编码器
	if global.App.Config.LogConf.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(), level)
}

// 使用 lumberjack 作为日志写入器
func getLogWriter() zapcore.WriteSyncer {
	timeStr := time.Now().Format("200601-02")
	file := &lumberjack.Logger{
		Filename:   global.App.Config.LogConf.RootDir + "/" + timeStr + ".log",
		MaxSize:    global.App.Config.LogConf.MaxSize,
		MaxBackups: global.App.Config.LogConf.MaxBackups,
		MaxAge:     global.App.Config.LogConf.MaxAge,
		Compress:   global.App.Config.LogConf.Compress,
	}

	return zapcore.AddSync(file)
}

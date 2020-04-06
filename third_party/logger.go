package third_party

import (
	"demoProject/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//初始化日志
func  InitLoger() (*zap.Logger,error) {
	log,err:=config.GetConfig()
	if err == nil {
		writeSyncer := getLogWriter(&log.LogConfig)
		encoder := getEncoder()
		core := zapcore.NewCore(encoder,writeSyncer,zapcore.ErrorLevel)
		logger := zap.New(core, zap.AddCaller())
		return logger,nil
	}
}

func  getEncoder()zapcore.Encoder  {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logConfig *config.LogConfig) zapcore.WriteSyncer {
	//增加日志文件切割和归档
	//file, _ := os.Create(path)
	//return zapcore.AddSync(file)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logConfig.Path,
		MaxSize:    logConfig.MaxSize,
		MaxBackups: logConfig.MaxBackups,
		MaxAge:     logConfig.MaxAge,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

package components

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() (core zapcore.Core) {
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey: "Message",
		LevelKey:   "Level",
		TimeKey:    "DateTime",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		CallerKey:    "File",
		EncodeCaller: zapcore.ShortCallerEncoder,
	})

	infoLog, _ := os.Create(conf.Log.InfoPath)
	errorLog, _ := os.Create(conf.Log.ErrorPath)

	return zapcore.NewTee(
		zapcore.NewCore(encoder, infoLog, zap.LevelEnablerFunc(func(z zapcore.Level) bool {
			return z >= zapcore.Level(conf.Log.Level) && z <= zap.WarnLevel
		})),
		zapcore.NewCore(encoder, errorLog, zap.LevelEnablerFunc(func(z zapcore.Level) bool {
			return z >= zap.ErrorLevel
		})),
	)
}

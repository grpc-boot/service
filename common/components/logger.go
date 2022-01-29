package components

import (
	"os"
	"time"

	"service/common/define/constant"

	"github.com/grpc-boot/base"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
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

	infoLog, _ := os.OpenFile(conf.Log.InfoPath, constant.LoggerFlag, constant.LoggerMode)
	errorLog, _ := os.OpenFile(conf.Log.ErrorPath, constant.LoggerFlag, constant.LoggerMode)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, infoLog, zap.LevelEnablerFunc(func(z zapcore.Level) bool {
			return z >= zapcore.Level(conf.Log.Level) && z <= zap.WarnLevel
		})),
		zapcore.NewCore(encoder, errorLog, zap.LevelEnablerFunc(func(z zapcore.Level) bool {
			return z >= zap.ErrorLevel
		})),
	)

	base.InitZapWithCore(core)
}

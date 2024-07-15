package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func init() {
	var config = zap.NewProductionEncoderConfig()

	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format(time.RFC3339))
	}
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder

	Logger = zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(config),
			zapcore.AddSync(os.Stderr),
			zapcore.InfoLevel,
		),
	).Sugar()
}

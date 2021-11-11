package zap

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"log"
	"time"
)

type FiberLogger struct {
	logger *zap.Logger
	Ctx *fiber.Ctx
}

func NewFiberLogger(Ctx *fiber.Ctx) Logger {
	if logger == nil {
		zaplog, err := zap.NewProduction()
		if err != nil {
			log.Printf("Error while create logger: %v",err)
		}
		defer zaplog.Sync()
		return &FiberLogger{
			logger: zaplog,
			Ctx: Ctx,
		}
	}
	return logger
}


func (f FiberLogger) Info(message string) {
	f.logger.Info(message,
		// Structured context as strongly typed Field values.
		zap.String("ClientIP", f.Ctx.IP()),
		zap.String("Proto", f.Ctx.Protocol()),
		zap.String("Method", f.Ctx.Method()),
		zap.String("Host", f.Ctx.Hostname()),
		zap.String("Path", f.Ctx.Path()),
		zap.ByteString("UserAgent", f.Ctx.Request().Header.UserAgent()),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func (f FiberLogger) Error(message string) {
	f.logger.Info(message,
		// Structured context as strongly typed Field values.
		zap.String("ClientIP", f.Ctx.IP()),
		zap.String("Proto", f.Ctx.Protocol()),
		zap.String("Method", f.Ctx.Method()),
		zap.String("Host", f.Ctx.Hostname()),
		zap.String("Path", f.Ctx.Path()),
		zap.ByteString("UserAgent", f.Ctx.Request().Header.UserAgent()),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func (f FiberLogger) Debug(message string) {
	panic("implement me")
}

func (f FiberLogger) Warning(message string) {
	panic("implement me")
}
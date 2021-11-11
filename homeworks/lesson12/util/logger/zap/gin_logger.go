package zap

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"time"
)


type Ginlogger struct {
	logger *zap.Logger
	Ctx *gin.Context
}

func NewGinLogger(Ctx *gin.Context) Logger {
	if logger == nil {
		zaplog, err := zap.NewProduction()
		if err != nil {
			log.Printf("Error while create logger: %v",err)
		}
		defer zaplog.Sync()
		return &Ginlogger{
			logger: zaplog,
			Ctx: Ctx,
		}
	}
	return logger
}


func (u Ginlogger) Info(message string) {
	u.logger.Info(message,
		// Structured context as strongly typed Field values.
		zap.String("ClientIP", u.Ctx.ClientIP()),
		zap.String("Proto", u.Ctx.Request.Proto),
		zap.String("Method", u.Ctx.Request.Method),
		zap.String("RemoteAddr",u.Ctx.Request.RemoteAddr),
		zap.String("Host", u.Ctx.Request.URL.Host),
		zap.String("Path", u.Ctx.Request.URL.Path),
		zap.String("UserAgent", u.Ctx.Request.UserAgent()),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func (u Ginlogger) Error(message string) {
	u.logger.Error(message,
		zap.String("ClientIP", u.Ctx.ClientIP()),
		zap.String("Proto", u.Ctx.Request.Proto),
		zap.String("Method", u.Ctx.Request.Method),
		zap.String("RemoteAddr",u.Ctx.Request.RemoteAddr),
		zap.String("Host", u.Ctx.Request.URL.Host),
		zap.String("Path", u.Ctx.Request.URL.Path),
		zap.String("UserAgent", u.Ctx.Request.UserAgent()),
		zap.Int("attempt", 5),
		zap.Duration("backoff", time.Second),
	)
}

func (u Ginlogger) Debug(message string) {
	panic("implement me")
}

func (u Ginlogger) Warning(message string) {
	panic("implement me")
}
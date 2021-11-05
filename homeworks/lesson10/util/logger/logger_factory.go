package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/kataras/iris/v12"
)

type FactoryLogger struct {
	GinContext *gin.Context
	IrisContext *iris.Context
	FiberContext *fiber.Ctx
}

func NewLogger(factory_logger FactoryLogger) Logger {
	if factory_logger.GinContext != nil {
		return NewGinLogger(factory_logger.GinContext)
	}
	if factory_logger.FiberContext != nil {
		return NewFiberLogger(factory_logger.FiberContext)
	}
	if factory_logger.IrisContext != nil {
		return NewGinLogger(factory_logger.GinContext)
	}
	return nil
}
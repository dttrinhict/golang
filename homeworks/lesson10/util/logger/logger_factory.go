package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/v12"
)

type FactoryLogger struct {
	GinContext *gin.Context
	IrisContext *iris.Context
}

func NewLogger(factory_logger FactoryLogger) Logger {
	if factory_logger.GinContext != nil {
		return NewGinLogger(factory_logger.GinContext)
	}
	if factory_logger.IrisContext != nil {
		return NewGinLogger(factory_logger.GinContext)
	}
	return nil
}
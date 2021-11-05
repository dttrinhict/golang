package zap

import (
	"go.uber.org/zap"
	"log"
	"time"
)

type MySQLLogger struct {
	logger *zap.Logger
}

var mysqlLogger Logger

func LoggerMysql() Logger {
	if mysqlLogger == nil {
		if zaplog == nil {
			zaplog, err := zap.NewProduction()
			if err != nil {
				log.Printf("Error while create logger: %v",err)
			}
			defer zaplog.Sync()
			return &MySQLLogger{
				logger: zaplog,
			}
		}
		return &MySQLLogger{
			logger: zaplog,
		}
	}
	return mysqlLogger
}


func (u MySQLLogger) Info(message string) {
	u.logger.Info(message,
		// Structured context as strongly typed Field values.
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func (u MySQLLogger) Error(message string) {
	u.logger.Error(message,
		zap.Int("attempt", 5),
		zap.Duration("backoff", time.Second),
	)
}

func (u MySQLLogger) Debug(message string) {
	panic("implement me")
}

func (u MySQLLogger) Warning(message string) {
	panic("implement me")
}
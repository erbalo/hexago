package log

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func NewLogger() logger.Option {
	logger.SetLogger()

	return logger.WithLogger(func(ctx *gin.Context, l zerolog.Logger) zerolog.Logger {
		return l.Output(gin.DefaultWriter).With().Logger()
	})
}

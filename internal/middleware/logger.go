package middleware

import (
	"gin-admin/global"
	"gin-admin/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		global.Logger.Info(path,
			logger.Int("status", c.Writer.Status()),
			logger.String("method", c.Request.Method),
			logger.String("path", path),
			logger.String("query", query),
			logger.String("ip", c.ClientIP()),
			logger.String("user-agent", c.Request.UserAgent()),
			logger.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			logger.Duration("cost", cost),
		)
	}
}

/**
 * @Author: LiuShuXin
 * @Description:
 * @File:  logger
 * Software: Goland
 * @Date: 2025/1/21 10:38
 */

package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// Logger 取代gin框架默认的日志中间件 记录请求信息到日志中
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求到达的时间
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		// 请求处理完成之后的一系列信息
		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nekvil/cars-go-api/internal/utils"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		var requestID uuid.UUID
		var err error
		requestID, err = uuid.NewV7()
		if err != nil {
			utils.Logger.Errorf("Failed to generate request ID: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Writer.Header().Set("X-Request-ID", requestID.String())

		c.Next()

		latency := time.Since(start)

		requestInfo := map[string]interface{}{
			"latency":    latency,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"clientIp":   c.ClientIP(),
			"statusCode": c.Writer.Status(),
			"requestId":  c.Writer.Header().Get("X-Request-ID"),
		}

		status := c.Writer.Status()

		var logLevel logrus.Level
		switch {
		case status >= http.StatusInternalServerError:
			logLevel = logrus.ErrorLevel
		case status >= http.StatusBadRequest:
			logLevel = logrus.WarnLevel
		default:
			logLevel = logrus.InfoLevel
		}

		utils.Logger.WithFields(requestInfo).Log(logLevel, "Request processed")
	}
}

package middleware

import (
	"time"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// RequestDurationMiddleware logs the duration of each request
func RequestDurationMiddleware(c *fiber.Ctx) error {
	startTime := time.Now()

	err := c.Next()

	duration := time.Since(startTime)
	requestID := c.Get("X-Request-Id")

	logger.Logger.Info(
		"Request completed",
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
		zap.Int("status_code", c.Response().StatusCode()),
		zap.Duration("duration_ms", duration),
		zap.String("request_id", requestID),
	)

	return err
}

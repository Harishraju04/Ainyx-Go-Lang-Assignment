package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// RequestIDMiddleware injects a unique requestId header in the response
func RequestIDMiddleware(c *fiber.Ctx) error {
	requestID := uuid.New().String()
	c.Set("X-Request-Id", requestID)
	c.Set("Access-Control-Expose-Headers", "X-Request-Id")
	return c.Next()
}

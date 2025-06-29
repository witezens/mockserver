package middleware

import (
	"time"

	"mock-server/utils"

	"github.com/gofiber/fiber/v2"
)

// RequestDurationLogger duration per each request
func RequestDurationLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		utils.Logger.Infow("Request completed",
			"method", c.Method(),
			"path", c.Path(),
			"status", c.Response().StatusCode(),
			"duration_us", duration.Microseconds(),
			"duration_ms", float64(duration.Microseconds())/1000.0,
		)

		return err
	}
}

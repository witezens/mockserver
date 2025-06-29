package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogRequest(c *fiber.Ctx, service, resource, mockFile string, status int, logBody bool) {
	logFields := []interface{}{
		"timestamp", time.Now().Format(time.RFC3339),
		"method", c.Method(),
		"path", c.Path(),
		"service", service,
		"resource", resource,
		"mock_file", mockFile,
		"status", status,
	}

	if logBody && (c.Method() == fiber.MethodPost || c.Method() == fiber.MethodPut) {
		var body map[string]interface{}
		_ = c.BodyParser(&body)
		if len(body) > 0 {
			logFields = append(logFields, "body", body)
		}
	}

	Logger.Infow("Request received", logFields...)
}

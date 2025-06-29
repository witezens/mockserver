package handler

import (
	"fmt"
	"mock-server/registry"
	"mock-server/resolver"
	"mock-server/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

func RegisterMockHandler(app *fiber.App, r *resolver.MockResolver) {
	handlers := registry.GetAll()

	app.All("/:service/api/v1/:resource", func(c *fiber.Ctx) error {
		service := c.Params("service")
		resource := c.Params("resource")
		method := c.Method()
		key := utils.BuildKey(service, resource)

		body := make(map[string]interface{})
		if method == fiber.MethodPost || method == fiber.MethodPut {
			_ = c.BodyParser(&body)
		}

		if handlerFunc, exists := handlers[key]; exists {
			return handlerFunc(c)
		}

		mockFile := r.ResolveFile(service, resource, method, body, utils.ToURLValues(c.Queries()))
		mockPath := fmt.Sprintf("mock-data/%s/api/v1/%s", service, mockFile)

		content, err := os.ReadFile(mockPath)
		if err != nil {
			utils.LogRequest(c, service, resource, mockFile, 404, false)
			return c.Status(404).JSON(fiber.Map{"error": "Mock not found"})
		}

		utils.LogRequest(c, service, resource, mockFile, 200, true)
		c.Type("application/json")
		return c.Send(content)
	})
}

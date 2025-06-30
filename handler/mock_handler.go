package handler

import (
	"mock-server/mockcache"
	"mock-server/registry"
	"mock-server/resolver"
	"mock-server/utils"

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
		if mockFile == "" {
			utils.LogRequest(c, service, resource, "", 404, true)
			return c.Status(404).JSON(fiber.Map{"error": "Mock not found"})
		}

		//  ¿In memory?
		if parsed, ok := mockcache.GlobalCache.Parsed[mockFile]; ok {
			utils.LogRequest(c, service, resource, mockFile, 200, true)
			return c.Status(200).JSON(parsed)
		}

		if raw, ok := mockcache.GlobalCache.Raw[mockFile]; ok {
			utils.LogRequest(c, service, resource, mockFile, 200, true)
			c.Type(utils.GetContentType(mockFile))
			return c.Send(raw)
		}

		// Nothing
		utils.LogRequest(c, service, resource, mockFile, 404, true)
		return c.Status(404).JSON(fiber.Map{"error": "Mock not found"})
	})
}

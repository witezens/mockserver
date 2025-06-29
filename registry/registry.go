package registry

import "github.com/gofiber/fiber/v2"

var handlers = map[string]fiber.Handler{}

func Register(route string, h fiber.Handler) {
	handlers[route] = h
}

func GetAll() map[string]fiber.Handler {
	return handlers
}

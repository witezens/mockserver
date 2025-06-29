package main

import (
	"mock-server/handler"
	"mock-server/middleware"
	"mock-server/resolver"
	"mock-server/utils"

	_ "mock-server/services/resourceinventory"

	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.InitLogger()
	defer utils.Logger.Sync()

	app := fiber.New()
	app.Use(middleware.RequestDurationLogger())

	myResolver := resolver.MockResolver{
		Rules: map[string][]resolver.MockRule{},
	}

	handler.RegisterMockHandler(app, &myResolver)

	err := app.Listen(":3000")

	if err != nil {
		return
	}
}

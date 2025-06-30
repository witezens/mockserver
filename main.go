package main

import (
	"log"
	"mock-server/handler"
	"mock-server/middleware"
	"mock-server/mockcache"
	"mock-server/resolver"
	"mock-server/utils"

	_ "mock-server/services/resourceinventory"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load mock files in memory
	if err := mockcache.GlobalCache.Load("mock-data"); err != nil {
		log.Fatalf("It cannot load mock files: %v", err)
	}

	utils.InitLogger()
	defer utils.Logger.Sync()

	app := fiber.New()
	app.Use(middleware.RequestDurationLogger())

	myResolver := resolver.MockResolver{
		Rules: map[string][]resolver.MockRule{
			"serviceinventory_ActiveServices": {
				{Param: "MSISDN", Source: "query", Versioned: true},
			},
			"executecollection_ExecuteCollection": {
				{Param: "collectionId", Source: "body", Versioned: false},
			},
		},
	}

	handler.RegisterMockHandler(app, &myResolver)

	err := app.Listen(":3000")

	if err != nil {
		return
	}
}

package resourceinventory

import (
	"fmt"
	"mock-server/mockcache"
	"mock-server/registry"
	"mock-server/utils"

	"github.com/gofiber/fiber/v2"
)

func HandleObtainReservedNumber(c *fiber.Ctx) error {
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
			"error": "Method Not Allowed",
		})
	}

	cachePath := "resourceinventory/api/v1/ObtainReservedNumber.__dynamic__.POST.json"

	template, ok := mockcache.GlobalCache.Parsed[cachePath]
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Template not found in cache: %s", cachePath),
		})
	}

	var body map[string]interface{}
	_ = c.BodyParser(&body)

	base := make(map[string]interface{})
	for k, v := range template {
		base[k] = v
	}

	// dynamic replaces
	if v, ok := body["ICC"]; ok {
		base["ICC"] = v
	} else {
		base["ICC"] = utils.RandomNumber(14)
	}

	if v, ok := body["IMSI"]; ok {
		base["IMSI"] = v
	} else {
		base["IMSI"] = utils.RandomNumber(10)
	}

	base["MSISDN"] = "2" + utils.RandomNumber(7)

	utils.LogRequest(c, "resourceinventory", "ObtainReservedNumber", cachePath, fiber.StatusOK, true)
	return c.Status(fiber.StatusOK).JSON(base)
}

func init() {
	registry.Register(utils.BuildKey("resourceinventory", "ObtainReservedNumber"), HandleObtainReservedNumber)
}

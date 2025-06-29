package resourceinventory

import (
	"encoding/json"
	"fmt"
	"mock-server/registry"
	"mock-server/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

func HandleObtainReservedNumber(c *fiber.Ctx) error {
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
			"error": "Method Not Allowed",
		})
	}

	// read base template
	path := "mock-data/resourceinventory/api/v1/ObtainReservedNumber.POST.json"
	content, err := os.ReadFile(path)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Cannot read template: %s", err),
		})
	}

	var body map[string]interface{}
	_ = c.BodyParser(&body)

	var base map[string]interface{}
	if err := json.Unmarshal(content, &base); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid template format",
		})
	}

	// ICC: 14 digits
	if v, ok := body["ICC"]; ok {
		base["ICC"] = v
	} else {
		base["ICC"] = utils.RandomNumber(14)
	}

	// IMSI: 10 digits
	if v, ok := body["IMSI"]; ok {
		base["IMSI"] = v
	} else {
		base["IMSI"] = utils.RandomNumber(10)
	}

	// MSISDN: init with 2 and contains 8 digits
	base["MSISDN"] = "2" + utils.RandomNumber(7)

	utils.LogRequest(c, "resourceinventory", "ObtainReservedNumber", path, fiber.StatusOK, true)
	return c.Status(fiber.StatusOK).JSON(base)
}

func init() {
	registry.Register(utils.BuildKey("resourceinventory", "ObtainReservedNumber"), HandleObtainReservedNumber)
}

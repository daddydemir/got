package handlers

import (
	"github.com/daddydemir/got/core"
	"github.com/daddydemir/got/models"
	"github.com/gofiber/fiber/v2"
)

func createModelHandler(c *fiber.Ctx) error {
	println(c.Method())
	if c.Method() == "POST" {
		var request models.ParseModel
		_ = c.BodyParser(&request)
		response := core.ParseWithJson(request.Data, "file")
		return c.SendString(response)
	} else {
		return c.SendString("get method")
	}
}

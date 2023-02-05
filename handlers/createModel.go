package handlers

import (
	"github.com/daddydemir/got/core"
	"github.com/daddydemir/got/models"
	"github.com/gofiber/fiber/v2"
)

func createModelHandler(c *fiber.Ctx) error {
	if c.Method() == "POST" {
		var request models.ParseModel
		_ = c.BodyParser(&request)
		core.ParseWithJson(request.Data)
		return c.Send([]byte(request.Data))
	} else {
		return c.SendString("get method")
	}
}

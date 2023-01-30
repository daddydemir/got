package handlers

import "github.com/gofiber/fiber/v2"

func createModelHandler(c *fiber.Ctx) error {
	if c.Method() == "POST" {
		return c.SendString("post method")
	} else {
		return c.SendString("get method")
	}
}

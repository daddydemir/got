package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func Urls() *fiber.App {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.All("/", dashHandler)
	return app
}

func dashHandler(c *fiber.Ctx) error {
	return c.Render("createModel", fiber.Map{})
}

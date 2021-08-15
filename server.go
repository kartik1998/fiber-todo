package main

import (
	"fiber-todo-poc/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "success",
			"result": "ApiOk",
		})
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "success",
			"result": "AoK",
		})
	})
	routes.TodoRoute(api.Group("/todos"))
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)
	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}

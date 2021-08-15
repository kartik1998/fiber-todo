package routes

import (
	"fiber-todo-poc/controllers"

	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
}

package routes

import (
	"fiber-todo-poc/controllers"

	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
	route.Post("", controllers.CreateTodo)
	route.Put("/:id", controllers.UpdateTodo)
	route.Delete("/:id", controllers.DeleteTodo)
	route.Get("/:id", controllers.GetTodo)
	route.Get("", controllers.GetTodos)
}

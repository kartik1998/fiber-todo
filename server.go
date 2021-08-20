package main

import (
	"fiber-todo-poc/controllers"
	"log"
	"os"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	if len(connectionString) == 0 {
		connectionString = "mongodb://localhost:27017"
	}
	err := mgm.SetDefaultConfig(nil, "todos", options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/api/todos", controllers.GetAllTodos)
	app.Get("/api/todos/:id", controllers.GetTodoByID)
	app.Post("/api/todos", controllers.CreateTodo)
	app.Patch("/api/todos/:id", controllers.ToggleTodoStatus)
	app.Delete("/api/todos/:id", controllers.DeleteTodo)

	app.Listen(":3000")
}

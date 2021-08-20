package controllers

import (
	"fiber-todo-poc/models"
	"net/http"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTodos(ctx *fiber.Ctx) error {
	collection := mgm.Coll(&models.Todo{})
	todos := []models.Todo{}

	err := collection.SimpleFind(&todos, bson.D{})
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return nil
	}

	ctx.JSON(fiber.Map{
		"ok":    true,
		"todos": todos,
	})
	return nil
}

func GetTodoByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return nil
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
	return nil
}

func CreateTodo(ctx *fiber.Ctx) error {
	params := new(struct {
		Title       string
		Description string
	})

	ctx.BodyParser(&params)

	if len(params.Title) == 0 || len(params.Description) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Title or description not specified.",
		})
		return nil
	}

	todo := models.CreateTodo(params.Title, params.Description)
	err := mgm.Coll(todo).Create(todo)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return nil
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
	return nil
}

func ToggleTodoStatus(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return nil
	}

	todo.Done = !todo.Done

	err = collection.Update(todo)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return nil
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
	return nil
}

func DeleteTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return nil
	}

	err = collection.Delete(todo)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return nil
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
	return nil
}

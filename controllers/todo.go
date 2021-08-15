package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{
	{
		Id:        1,
		Title:     "Walk the dog",
		Completed: false,
	},
	{
		Id:        2,
		Title:     "Walk the cat",
		Completed: false,
	},
}

func GetTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"result": fiber.Map{
			"todos": todos,
		},
	})
}

func CreateTodo(c *fiber.Ctx) error {
	type Request struct {
		Title string `json:"title"`
	}
	var body Request
	err := c.BodyParser(&body)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Cannot parse JSON",
		})
	}

	// create a todo variable
	todo := &Todo{
		Id:        len(todos) + 1,
		Title:     body.Title,
		Completed: false,
	}

	// append in todos
	todos = append(todos, todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"result": fiber.Map{
			"todo": todo,
		},
	})
}

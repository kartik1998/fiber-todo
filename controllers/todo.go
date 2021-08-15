package controllers

import (
	"fmt"
	"strconv"

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

func GetTodo(c *fiber.Ctx) error {
	// get parameter value
	paramId := c.Params("id")

	// convert parameter value string to int
	id, err := strconv.Atoi(paramId)

	// if error in parsing string to int
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Cannot parse Id",
		})
	}

	// find todo and return
	for _, todo := range todos {
		if todo.Id == id {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status": "success",
				"result": fiber.Map{
					"todo": todo,
				},
			})
		}
	}

	// if todo not available
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"status": "error",
		"error":  "Todo not found",
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	// find parameter
	paramId := c.Params("id")

	// convert parameter string to int
	id, err := strconv.Atoi(paramId)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Cannot parse id",
		})
	}

	// request structure
	type Request struct {
		Title     *string `json:"title"`
		Completed *bool   `json:"completed"`
	}

	var body Request
	err = c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"result": "Cannot parse JSON",
		})
	}

	var todo *Todo

	for _, t := range todos {
		if t.Id == id {
			todo = t
			break
		}
	}

	if todo.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  "Not found",
		})
	}

	if body.Title != nil {
		todo.Title = *body.Title
	}

	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"result": fiber.Map{
			"todo": todo,
		},
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	// get param
	paramId := c.Params("id")

	// convert param string to int
	id, err := strconv.Atoi(paramId)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Cannot parse id",
		})
	}

	// find and delete todo
	for i, todo := range todos {
		if todo.Id == id {

			todos = append(todos[:i], todos[i+1:]...)

			return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
				"status": "success",
				"result": "Deleted Succesfully",
			})
		}
	}

	// if todo not found
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"status": "error",
		"error":  "Todo not found",
	})
}

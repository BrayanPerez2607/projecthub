package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BrayanPerez2607/projecthub/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	todos := []models.ToDo{}

	//Get all the ToDo's
	app.Get("/api/atodos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	//Create a ToDo
	app.Post("/api/ctodos", func(c *fiber.Ctx) error {
		todo := &models.ToDo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "ToDo can't be empty"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	//Update a ToDo
	app.Patch("/api/utodos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(400).JSON(fiber.Map{"error": "ToDo not found"})
	})

	//Delete a ToDo
	app.Delete("/api/dtodos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": "true"})
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "ToDo not found"})
	})

	log.Fatal(app.Listen(":" + port))
}

package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}


func main() {
	fmt.Println("Hello worlds")
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	todos := []Todo{} //Array
	
	//Get Todos
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(201).JSON(todos) 
	})

	//Create a Todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{} // If empty, it has false values: {id: 0, completed:false,body:""}

		if err := c.BodyParser(todo); err != nil {
			return err
		}
	 
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	//Update a Todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	//Delete a Todo
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
					// SAMPLE
					// 1 2 3 4 5
					// todos[:i] RESULT: 1 2 
					// todos[i+1:]... get the todos after the inputted todo not including the inputted one RESULT: 4 5
					// FINAL: 1 2 4 5
				return c.Status(200).JSON(fiber.Map{"success": true})
			} 
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})
	

	log.Fatal(app.Listen(":"+PORT))
}
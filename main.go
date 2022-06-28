package main

import (
	"Golang_API/app"
	"Golang_API/configs"
	"Golang_API/repository"
	"Golang_API/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appRoute := fiber.New()

	configs.ConnectDB()

	dbClient := configs.GetCollection(configs.DB, "todos")

	TodoRepositoryDb := repository.NewTodoRepositoryDb(dbClient)

	td := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDb)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAll)
	appRoute.Delete("/api/todo/:id", td.Delete)

	appRoute.Listen(":8080")

}

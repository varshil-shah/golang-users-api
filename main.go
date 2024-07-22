package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/varshil-shah/golang-users-api/configs"
	"github.com/varshil-shah/golang-users-api/routers"
)

func main() {
	fmt.Println("Welcome to API using GoLang, GoFiber & MongoDB!")

	app := fiber.New()

	configs.ConnectDatabase()

	routers.UserRouter(app)

	app.Get("/", func(context *fiber.Ctx) error {
		return context.JSON(&fiber.Map{"data": "Hello from GoFiber & MongoDB"})
	})

	app.Listen(":3000")
}

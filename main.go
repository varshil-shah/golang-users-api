package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/varshil-shah/golang-users-api/configs"
	"github.com/varshil-shah/golang-users-api/routers"
)

func main() {
	fmt.Println("Welcome to API using GoLang, GoFiber & MongoDB!")

	app := fiber.New()

	app.Use(cors.New())

	configs.ConnectDatabase()

	routers.UserRouter(app)

	app.Get("/", func(context *fiber.Ctx) error {
		return context.JSON(&fiber.Map{"data": "Hello from GoFiber & MongoDB"})
	})

	defer func() {
		fmt.Println("MongoDB connection disconnected!")
		configs.DB.Disconnect(context.Background())
	}()

	port, _ := configs.GetEnvVariable("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen("0.0.0.0" + port)
}

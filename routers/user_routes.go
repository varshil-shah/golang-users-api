package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/varshil-shah/golang-users-api/controllers"
)

func UserRouter(app *fiber.App) {
	app.Get("/users", controllers.GetAllUsers)
	app.Post("/users", controllers.CreateUser)

	app.Get("/users/:userId", controllers.GetUser)
	app.Put("/users/:userId", controllers.UpdateUser)
	app.Delete("/users/:userId", controllers.DeleteUser)
}

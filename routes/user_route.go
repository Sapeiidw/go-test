package routes

import (
	"go-test/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, controller *controllers.UserController) {
	userRoutes := app.Group("/users")
	userRoutes.Get("/", controller.GetUsers)
	userRoutes.Get("/:id", controller.GetUser)
	userRoutes.Post("/", controller.CreateUser)
	userRoutes.Put("/:id", controller.UpdateUser)
	userRoutes.Delete("/:id", controller.DeleteUser)
}

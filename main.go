package main

import (
	"go-test/config"
	"go-test/controllers"
	"go-test/middleware"
	"go-test/repository"
	"go-test/routes"
	"go-test/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.ConnectDatabase()

	app := fiber.New()

	// Apply logging middleware
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true, // Show stack trace in logs (useful for debugging)
	})) // ✅ This will catch any panics and prevent crashes
	app.Use(middleware.LoggerMiddleware())


	// Setup Dependencies
	userRepo := repository.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.SetupUserRoutes(app, userController)



	app.Listen(":3000")
}

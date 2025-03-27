package main

import (
	"go-test/config"
	"go-test/controllers"
	"go-test/repository"
	"go-test/routes"
	"go-test/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDatabase()

	app := fiber.New()

	userRepo := repository.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.SetupUserRoutes(app, userController)

	app.Listen(":3000")
}

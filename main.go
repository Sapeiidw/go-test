package main

import (
	"go-test/config"
	"go-test/controllers"
	"go-test/middleware"
	"go-test/repository"
	"go-test/routes"
	"go-test/services"
	"log"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
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


	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open("./input.html")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		log.Fatal(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile("./output.pdf")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done")

	app.Listen(":3000")
}

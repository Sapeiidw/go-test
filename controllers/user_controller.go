package controllers

import (
	"go-test/models"
	"go-test/services"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) GetUsers(ctx *fiber.Ctx) error {
	users, err := c.service.GetAllUsers()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Failed to fetch users"})
	}
	return ctx.JSON(users)
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	user, err := c.service.GetUserByID(uint(id))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return ctx.JSON(user)
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	err := c.service.CreateUser(user)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}
	return ctx.JSON(user)
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	user, err := c.service.GetUserByID(uint(id))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	c.service.UpdateUser(user)
	return ctx.JSON(user)
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	// Parse and validate ID
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	userID := uint(id) // Convert once and reuse

	// Check if user exists before deleting
	if _, err := c.service.GetUserByID(userID); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Delete the user
	if err := c.service.DeleteUser(userID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})
}
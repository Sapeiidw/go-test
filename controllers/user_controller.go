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
	id, _ := ctx.ParamsInt("id")
	err := c.service.DeleteUser(uint(id))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Failed to delete user"})
	}
	return ctx.JSON(fiber.Map{"message": "User deleted successfully"})
}

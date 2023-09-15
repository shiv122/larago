package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shiv122/go-test/app/repositories"
)

type UserController struct{}

func (uc *UserController) GetUsers(c *fiber.Ctx) error {

	var userRepo = new(repositories.UserRepository)

	users, err := userRepo.GetPaginatedUsers(c.QueryInt("page", 1), c.QueryInt("pageSize", 5))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(users)
}

func (uc *UserController) GetUserById(c *fiber.Ctx) error {
	// Get user ID from the request parameters
	// Fetch the user from the database
	// Serialize it into JSON
	user := "test"

	return c.JSON(user)
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	// Parse request body to get user data
	// Create user in the database
	// Return success message
	return c.SendString("User created successfully")
}

func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	// Get user ID from the request parameters
	// Parse request body to get updated user data
	// Update user in the database
	// Return success message
	return c.SendString("User updated successfully")
}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	// Get user ID from the request parameters
	// Delete user from the database
	// Return success message
	return c.SendString("User deleted successfully")
}

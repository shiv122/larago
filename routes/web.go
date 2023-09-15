package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shiv122/go-test/app/controllers"
)

func SetupRoutes(route *fiber.App) {

	userController := new(controllers.UserController)

	route.Get("/users", userController.GetUsers)

	route.Get("/users/:id", userController.GetUserById)

	route.Post("/users", userController.CreateUser)

	route.Put("/users/:id", userController.UpdateUser)
	
	route.Delete("/users/:id", userController.DeleteUser)
}

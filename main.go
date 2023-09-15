package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/shiv122/go-test/config"
	"github.com/shiv122/go-test/database"
	"github.com/shiv122/go-test/routes"
)

func main() {
	config.LoadConfig()

	app := fiber.New()

	app.Use(pprof.New())

	app.Use(pprof.New(pprof.Config{Prefix: "/debug/pprof/"}))
	// Initialize and connect to the database
	config.ConnectDB()

	// Run database migrations
	database.Migrate()

	// Setup routes
	routes.SetupRoutes(app)

	// Start the Fiber server
	err := app.Listen(":" + config.Port)
	if err != nil {
		panic(err)
	}
}

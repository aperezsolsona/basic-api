package main

import (
	"github.com/aperezsolsona/basic-api/database"
	"github.com/aperezsolsona/basic-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb(true)
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/user", handlers.ListUsers)
	app.Post("/user", handlers.CreateUser)
	// Add new route to show single User, given `:id`
	app.Get("/user/:id", handlers.ShowUser)
}

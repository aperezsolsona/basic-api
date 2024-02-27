package main

import (
	"github.com/aperezsolsona/basic-api/database"
	"github.com/aperezsolsona/basic-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb(true, false)
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Get("/users", handlers.ListUsers)
	app.Get("/users/:id", handlers.ShowUser)
	app.Post("/users", handlers.CreateUser)

	app.Get("/pets", handlers.ListPets)
	app.Get("/pets/:id", handlers.ShowPet)
	app.Post("/pets", handlers.CreatePet)

	app.Get("/pet-types", handlers.ListPetTypes)
}

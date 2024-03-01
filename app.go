package main

import (
	"github.com/aperezsolsona/basic-api/database"
	"github.com/aperezsolsona/basic-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect(true, false)
	app := fiber.New()
	SetupRoutes(app)
	app.Listen(":3000")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", Ping)

	app.Get("/users", handlers.ListUsers)
	app.Get("/users/:id", handlers.ShowUser)
	app.Post("/users", handlers.CreateUser)

	app.Get("/pets", handlers.ListPets)
	app.Get("/pets/:id", handlers.ShowPet)
	app.Post("/pets", handlers.CreatePet)

	app.Get("/pet-types", handlers.ListPetTypes)
}

func Ping(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "Basic-API is running!"})
}

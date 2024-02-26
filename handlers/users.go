package handlers

import (
	"github.com/aperezsolsona/basic-api/database"
	"github.com/aperezsolsona/basic-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UserResponse represents the api response representation of a user
type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func Home(c *fiber.Ctx) error {
	// TODO openAPI
	return c.SendString("This should be some sort of OpenAPI!")
}

func ListUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Db.Find(&users)

	return c.Status(200).JSON(mapUserResponses(users))
}

func ShowUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id") // from URL
	database.DB.Db.Where("id = ?", id).First(&user)

	return c.Status(200).JSON(mapUserResponse(user))
}

func CreateUser(c *fiber.Ctx) error {
	// Parse request body struct
	type RequestBody struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var requestBody RequestBody
	// Parse the fiber context
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Generate UUID for the user
	userID := uuid.New()

	// Create User object
	user := models.User{
		ID:    userID,
		Name:  requestBody.Name,
		Email: requestBody.Email,
	}

	database.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}

func mapUserResponse(user models.User) UserResponse {

	userResponse := UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return userResponse
}

func mapUserResponses(users []models.User) []UserResponse {

	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, mapUserResponse(user))
	}

	return userResponses
}

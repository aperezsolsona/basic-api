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

func ListUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.DbInstance.Db.Find(&users)

	return c.Status(200).JSON(MapResponses(users))
}

func ShowUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id") // from URL
	database.DbInstance.Db.Where("id = ?", id).First(&user)

	return c.Status(200).JSON(MapResponse(user))
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

	database.DbInstance.Db.Create(&user)

	return c.Status(200).JSON(user)
}

func MapResponse(user models.User) UserResponse {

	userResponse := UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return userResponse
}

func MapResponses(users []models.User) []UserResponse {

	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, MapResponse(user))
	}

	return userResponses
}

package handlers

import (
	"github.com/aperezsolsona/basic-api/database"
	"github.com/aperezsolsona/basic-api/models"
	"github.com/gofiber/fiber/v2"
)

// PetTypeResponse represents the api response representation of a pettype
type PetTypeResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ListPetTypes(c *fiber.Ctx) error {
	petTypes := []models.PetType{}
	database.DbInstance.Db.Find(&petTypes)

	return c.Status(200).JSON(MapResponses(petTypes))
}

func MapResponse(petType models.PetType) PetTypeResponse {

	petTypeResponse := PetTypeResponse{
		ID:   petType.ID,
		Name: petType.Name,
	}

	return petTypeResponse
}

func MapResponses(petTypes []models.PetType) []PetTypeResponse {

	var petTypeResponses []PetTypeResponse
	for _, petType := range petTypes {
		petTypeResponses = append(petTypeResponses, MapResponse(petType))
	}

	return petTypeResponses
}

package handlers

import (
	"github.com/aperezsolsona/basic-api/database"
	"github.com/aperezsolsona/basic-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PetResponse represents the api response representation of a pet
type PetResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	PetTypeName string    `json:"pet_type_name"`
}

func ListPets(c *fiber.Ctx) error {
	pets := []models.Pet{}
	database.DB.Db.Preload("PetType").Find(&pets)

	return c.Status(200).JSON(mapPetResponses(pets))
}

func ShowPet(c *fiber.Ctx) error {
	pet := models.Pet{}
	id := c.Params("id") // from URL
	database.DB.Db.Where("id = ?", id).Preload("PetType").First(&pet)

	return c.Status(200).JSON(mapPetResponse(pet))
}

func CreatePet(c *fiber.Ctx) error {
	// Parse request body struct
	type RequestBody struct {
		Name      string `json:"name"`
		PetTypeID uint   `json:"pet_type_id"`
	}

	// Validate body
	var requestBody RequestBody
	// Parse the fiber context
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Validate Pet Type ID
	petType := models.PetType{}
	result := database.DB.Db.Where("id = ?", requestBody.PetTypeID).First(&petType)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Pet Type does not exist!",
		})
	}

	// Generate UUID for the pet
	petID := uuid.New()

	// Create Pet object
	pet := models.Pet{
		ID:        petID,
		Name:      requestBody.Name,
		PetTypeID: requestBody.PetTypeID,
	}

	database.DB.Db.Create(&pet)

	return c.Status(200).JSON(pet)
}

func mapPetResponse(pet models.Pet) PetResponse {

	petResponse := PetResponse{
		ID:          pet.ID,
		Name:        pet.Name,
		PetTypeName: pet.PetType.Name,
	}

	return petResponse
}

func mapPetResponses(pets []models.Pet) []PetResponse {

	var petResponses []PetResponse
	for _, pet := range pets {
		petResponses = append(petResponses, mapPetResponse(pet))
	}

	return petResponses
}

package handlers

import (
	"testing"

	"github.com/aperezsolsona/basic-api/models"
	"github.com/google/uuid"
)

func TestMapUserResponse(t *testing.T) {
	testuuid := uuid.New()
	testuser := models.User{
		ID:    testuuid,
		Name:  "Alex",
		Email: "a@a.es",
	}

	result := mapUserResponse(testuser)

	userResponseExpected := UserResponse{
		ID:    testuuid,
		Name:  "Alex",
		Email: "a@a.es",
	}

	if result != userResponseExpected {
		t.Errorf("errorrrrrr")
		//t.Errorf("mapping returned %d, expected %d", result, userResponseExpected)
	}
}

package database

import (
	"fmt"
	"log"
	"os"

	"github.com/aperezsolsona/basic-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb(isMigration bool) {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Madrid",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("DB connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	if isMigration {
		log.Println("Running DB migrations")
		db.AutoMigrate(&models.User{}, &models.PetType{}, &models.Pet{})

		// Create pet types
		petType1 := models.PetType{Name: "Dog"}
		petType2 := models.PetType{Name: "Cat"}
		db.Create(&petType1)
		db.Create(&petType2)
	}

	DB = Dbinstance{
		Db: db,
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/onainadapdap1/jwt-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dataSources struct {
	DB *gorm.DB
}

func initDS() (*dataSources, error) {
	log.Printf("Initializing data sources\n")
	dbUrl := os.Getenv("DATABASE_URL")

	log.Printf("connecting to postgresql\n")
	db, err := gorm.Open(postgres.Open(dbUrl))

	if err != nil {
		return nil, fmt.Errorf("error opening db : %w", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, fmt.Errorf("error migrating models: %w", err)
	}

	return &dataSources{DB: db}, nil
}
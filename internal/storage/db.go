package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect DataBase initialise the connection to the database
func ConnectDataBase() {
	dsn := "host=localhost user=postgres password=Jaimatadi@2020 dbname=ez_apply port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")
	DB = db

}

package database

import (
	"fmt"
	"log"
	"strconv"

	"fiber-demo/config"
	"fiber-demo/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	port, err := strconv.ParseUint(config.Config("DB_PORT"), 10, 32)

	if err != nil {
		log.Println("bad port")
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

    fmt.Printf("Connection Opened to db: %s\n", config.Config("DB_NAME"))

    // Migrate the db
    db.AutoMigrate(&model.Note{})
    fmt.Printf("Database Migrated\n")
}

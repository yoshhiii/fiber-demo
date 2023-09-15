package database

import (
	"fmt"
	"log"
	"strconv"

	"fiber-demo/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDB() {
	var err error
	port, err := strconv.ParseUint(config.Config("DB_PORT"), 10, 32)

	if err != nil {
		log.Println("bad port")
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"))

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// check if database exists
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", config.Config("DB_NAME"))
	rs := db.Raw(stmt)

	if rs.Error != nil {
		fmt.Println("Database does not exist")
	}

	// create database if it does not exist
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", config.Config("DB_NAME"))
		if rs := db.Exec(stmt); rs.Error != nil {
			fmt.Println("Error creating database")
		}

		//close db Connection
		sql, err := db.DB()
		defer func() {
			_ = sql.Close()
		}()

		if err != nil {
			fmt.Println("Error closing database")
		}
	}

    fmt.Printf("%s created or exists", config.Config("DB_NAME"))
}

package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewConnection() (*gorm.DB, error) {
	dsn := "host=auth-bd user=notification password=8QSDGT9H dbname=notification port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

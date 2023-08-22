package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DataBase() (*gorm.DB, error) {
	dbUri := "host=localhost port=5432 user=postgres password=******** dbname=book sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return db, nil
}

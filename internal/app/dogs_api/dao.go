package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDao(dbUrl string) (db *gorm.DB) {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

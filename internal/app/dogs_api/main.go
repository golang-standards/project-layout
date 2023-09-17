package main

import (
	"github.com/TriplePi/dogs/internal/app/dogs_api/domains/dog"
	user "github.com/TriplePi/dogs/internal/app/dogs_api/domains/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbUrl := "postgres://sobaka:nikitos_sosaka@localhost:5432/sobaki?sslmode=disable"
	RunMigrations(dbUrl)
	db := InitDao(dbUrl)

	router := gin.Default()
	dog.InitDogHandlers(router, db)
	router.Run("localhost:8080")
}

func InitDao(dbUrl string) (db *gorm.DB) {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&user.UserRole{})
	db.AutoMigrate(&dog.Dog{})
	return db
}

package config

import (
	"github.com/TriplePi/dogs/internal/app/dogs_api/domains/dog"
	"github.com/TriplePi/dogs/internal/app/dogs_api/domains/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb(dbUrl string) *gorm.DB {
	RunMigrations(dbUrl)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	_ = db.AutoMigrate(&user.User{})
	_ = db.AutoMigrate(&user.UserRole{})
	_ = db.AutoMigrate(&dog.Dog{})

	return db
}

func initRouter(db *gorm.DB) {
	router := gin.Default()
	dog.InitDogHandlers(router, db)
	user.InitUserHandlers(router, db)
	_ = router.Run("localhost:8080")
}

func Init() {
	dbUrl := "postgres://sobaka:nikitos_sosaka@localhost:5432/sobaki?sslmode=disable"

	db := initDb(dbUrl)
	initRouter(db)
}

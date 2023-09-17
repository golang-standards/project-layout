package config

import (
	"fmt"
	"github.com/TriplePi/dogs/internal/app/dogs_api/domains/dog"
	"github.com/TriplePi/dogs/internal/app/dogs_api/domains/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var envPath = "configs/.env"

type Environment struct {
	db  DbConfig
	app AppConfig
}

type DbConfig struct {
	url      string
	name     string
	login    string
	password string
}

type AppConfig struct {
	host string
	port string
}

func initDb(environment Environment) *gorm.DB {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		environment.db.login, environment.db.password, environment.db.url, environment.db.name,
	)

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

func initRouter(db *gorm.DB, environment Environment) {
	router := gin.Default()
	dog.InitDogHandlers(router, db)
	user.InitUserHandlers(router, db)
	_ = router.Run(environment.app.host + ":" + environment.app.port)
}

func Init() {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var environment = Environment{
		app: AppConfig{
			host: os.Getenv("APP_HOST"),
			port: os.Getenv("APP_PORT"),
		},
		db: DbConfig{
			url:      os.Getenv("DB_URL"),
			name:     os.Getenv("DB_NAME"),
			login:    os.Getenv("DB_LOGIN"),
			password: os.Getenv("DB_PASSWORD"),
		},
	}

	db := initDb(environment)
	initRouter(db, environment)
}

package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
	"time"
)

type Dog struct {
	gorm.Model
	Tattoo                 string    `json:"tattoo,omitempty"`
	ChipNumber             int64     `json:"chip_number,omitempty"`
	Name                   string    `json:"name,omitempty"`
	BirthDate              time.Time `json:"birth_date"`
	Breed                  string    `json:"breed,omitempty"`
	Sex                    bool      `json:"sex,omitempty"`
	Owner                  User      `json:"owner,omitempty" gorm:"foreignKey:OwnerID"`
	OwnerID                uint      `json:"owner_id,omitempty"`
	Coach                  User      `json:"coach,omitempty" gorm:"foreignKey:CoachID"`
	CoachID                uint      `json:"coach_id,omitempty" `
	BookForSportsmenNumber string    `json:"book_for_sportsmen_number,omitempty"`
}

type User struct {
	gorm.Model
	Username   string     `json:"username,omitempty"`
	Firstname  string     `json:"firstname,omitempty"`
	Lastname   string     `json:"lastname,omitempty"`
	Patronymic string     `json:"patronymic,omitempty"`
	Roles      []UserRole `json:"roles,omitempty"`
}

type UserRole struct {
	gorm.Model
	UserID uint
	Role   Role `gorm:"type:role"`
}

type Role string

const (
	OWNER = "OWNER"
	COACH = "COACH"
	JUDGE = "JUDGE"
	DECOY = "DECOY"
)

var users = []User{
	{
		Username:   "triplepi",
		Firstname:  "Pavel",
		Lastname:   "Popov",
		Patronymic: "Pavlovi4",
	},
}

var dogs = []Dog{
	{
		Tattoo:                 "Roma",
		ChipNumber:             666,
		Name:                   "Roma",
		BirthDate:              time.Time{},
		Breed:                  "Pekingese",
		Sex:                    true,
		Owner:                  users[0],
		Coach:                  users[0],
		BookForSportsmenNumber: "666",
	},
}

var globalDb gorm.DB

func main() {
	dbUrl := "postgres://sobaka:nikitos_sosaka@localhost:5432/sobaki?sslmode=disable"
	log.Print("db init started")
	RunMigrations(dbUrl)
	db := InitDao(dbUrl)
	globalDb = *db
	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserRole{})
	db.AutoMigrate(&Dog{})

	//db.Create(&users[0])
	//var user User
	//db.First(&user)
	//dogs[0].Owner = user
	//dogs[0].Coach = user
	//db.Create(&dogs[0])

	log.Print("db init finished")
	var dog []Dog
	db.Preload(clause.Associations).Find(&dog)
	//db.Find(&dog, "1")
	router := gin.Default()
	router.GET("/dogs", GetDogs)
	router.GET("/dogs/:id", GetDogsById)
	router.POST("/dogs", CreateDog)
	router.DELETE("/dogs/:id", DeleteDog)
	router.PATCH("/dogs/:id", UpdateDog)
	router.Run("localhost:8080")
}

func GetDogs(c *gin.Context) {
	var dogs []Dog
	globalDb.Preload(clause.Associations).Find(&dogs)
	c.IndentedJSON(http.StatusOK, dogs)
}

func GetDogsById(c *gin.Context) {
	var dog Dog
	err := globalDb.Preload(clause.Associations).First(&dog, c.Param("id")).Error
	if err != nil && err.Error() == "record not found" {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.IndentedJSON(http.StatusOK, dog)
}

func CreateDog(c *gin.Context) {
	var dog Dog
	if err := c.BindJSON(&dog); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	globalDb.Create(&dog)
	c.IndentedJSON(http.StatusCreated, dog)
}

func DeleteDog(c *gin.Context) {
	id := c.Param("id")
	globalDb.Delete(&Dog{}, id)
}

func UpdateDog(c *gin.Context) {
	id := c.Param("id")
	var storedDog Dog
	var patchedDog Dog
	err := globalDb.Preload(clause.Associations).First(&storedDog, id).Error
	if err != nil && err.Error() == "record not found" {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	if err := c.BindJSON(&patchedDog); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

}

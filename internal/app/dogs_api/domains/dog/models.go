package dog

import (
	user "github.com/TriplePi/dogs/internal/app/dogs_api/domains/user"
	"gorm.io/gorm"
	"time"
)

type Dog struct {
	gorm.Model
	Tattoo                 string    `json:"tattoo,omitempty"`
	ChipNumber             int64     `json:"chipNumber,omitempty"`
	Name                   string    `json:"name,omitempty"`
	BirthDate              time.Time `json:"birthDate"`
	Breed                  string    `json:"breed,omitempty"`
	Sex                    bool      `json:"sex,omitempty"`
	Owner                  user.User `json:"owner,omitempty" gorm:"foreignKey:OwnerID;references:ID"`
	OwnerID                uint      `json:"-"`
	Coach                  user.User `json:"coach,omitempty" gorm:"foreignKey:CoachID;references:ID"`
	CoachID                uint      `json:"-"`
	BookForSportsmenNumber string    `json:"bookForSportsmenNumber,omitempty"`
}

type UpdateDto struct {
	Tattoo                 *string    `json:"tattoo,omitempty"`
	ChipNumber             *int64     `json:"chipNumber,omitempty"`
	Name                   *string    `json:"name,omitempty"`
	BirthDate              *time.Time `json:"birthDate"`
	Breed                  *string    `json:"breed,omitempty"`
	Sex                    *bool      `json:"sex,omitempty"`
	OwnerID                *uint      `json:"ownerId,omitempty"`
	CoachID                *uint      `json:"coachId,omitempty"`
	BookForSportsmenNumber *string    `json:"bookForSportsmenNumber,omitempty"`
}

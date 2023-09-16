package dog

import (
	"gorm.io/gorm"
	"os/user"
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
	Owner                  user.User `json:"owner,omitempty" gorm:"foreignKey:OwnerID"`
	OwnerID                uint      `json:"ownerId,omitempty"`
	Coach                  user.User `json:"coach,omitempty" gorm:"foreignKey:CoachID"`
	CoachID                uint      `json:"coachId,omitempty" `
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

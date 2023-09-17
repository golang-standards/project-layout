package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string     `json:"username,omitempty"`
	Firstname  string     `json:"firstname,omitempty"`
	Lastname   string     `json:"lastname,omitempty"`
	Patronymic string     `json:"patronymic,omitempty"`
	Roles      []UserRole `json:"roles,omitempty"`
}

type UpdateDto struct {
	Username   *string     `json:"username,omitempty"`
	Firstname  *string     `json:"firstname,omitempty"`
	Lastname   *string     `json:"lastname,omitempty"`
	Patronymic *string     `json:"patronymic,omitempty"`
	Roles      *[]UserRole `json:"roles,omitempty"`
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

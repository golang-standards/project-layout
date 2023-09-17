package user

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Service struct {
	db *gorm.DB
}

func InitUserService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) GetAll() []User {
	var users []User
	s.db.Preload(clause.Associations).Find(&users)
	return users
}

func (s *Service) GetById(id string) (*User, error) {
	var user User
	err := s.db.Preload(clause.Associations).First(&user, id).Error
	if err != nil && err.Error() == "record not found" {
		return nil, err
	}
	return &user, nil
}

func (s *Service) Create(user *User) {
	s.db.Create(user)
}

func (s *Service) Delete(id string) {
	s.db.Delete(&User{}, id)
}

func (s *Service) Update(id string, dto *UpdateDto) (*User, error) {
	var storedUser User
	err := s.db.Preload(clause.Associations).First(&storedUser, id).Error
	if err != nil && err.Error() == "record not found" {
		return nil, err
	}
	var updatedUser = dto.ToEntity(storedUser)
	s.db.Updates(&updatedUser)
	return &updatedUser, nil
}

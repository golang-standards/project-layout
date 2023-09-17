package dog

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Service struct {
	db *gorm.DB
}

func InitDogService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) GetDogs() []Dog {
	var dogs []Dog
	s.db.Preload(clause.Associations).Find(&dogs)
	return dogs
}

func (s *Service) GetDogById(id string) (*Dog, error) {
	var dog Dog
	err := s.db.Preload(clause.Associations).First(&dog, id).Error
	if err != nil && err.Error() == "record not found" {
		return nil, err
	}
	return &dog, nil
}

func (s *Service) CreateDog(dog *Dog) {
	s.db.Create(dog)
}

func (s *Service) DeleteDog(id string) {
	s.db.Delete(&Dog{}, id)
}

func (s *Service) UpdateDog(id string, dto *UpdateDto) (*Dog, error) {
	var storedDog Dog
	err := s.db.Preload(clause.Associations).First(&storedDog, id).Error
	if err != nil && err.Error() == "record not found" {
		return nil, err
	}
	var updatedDog = dto.ToEntity(storedDog)
	s.db.Updates(&updatedDog)
	return &updatedDog, nil
}

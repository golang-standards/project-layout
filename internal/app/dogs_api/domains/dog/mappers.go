package dog

import "github.com/TriplePi/dogs/internal/app/dogs_api/utils"

func (dto UpdateDto) ToEntity(dog Dog) Dog {
	utils.ReplaceIfNotNil(dto.Tattoo, &dog.Tattoo)
	utils.ReplaceIfNotNil(dto.ChipNumber, &dog.ChipNumber)
	utils.ReplaceIfNotNil(dto.Name, &dog.Name)
	utils.ReplaceIfNotNil(dto.BirthDate, &dog.BirthDate)
	utils.ReplaceIfNotNil(dto.Breed, &dog.Breed)
	utils.ReplaceIfNotNil(dto.Sex, &dog.Sex)
	utils.ReplaceIfNotNil(dto.OwnerID, &dog.OwnerID)
	utils.ReplaceIfNotNil(dto.CoachID, &dog.CoachID)
	utils.ReplaceIfNotNil(dto.BookForSportsmenNumber, &dog.BookForSportsmenNumber)
	return dog
}

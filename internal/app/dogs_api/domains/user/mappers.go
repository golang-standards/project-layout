package user

import "github.com/TriplePi/dogs/internal/app/dogs_api/utils"

func (dto UpdateDto) ToEntity(entity User) User {
	utils.ReplaceIfNotNil(dto.Username, &entity.Username)
	utils.ReplaceIfNotNil(dto.Firstname, &entity.Firstname)
	utils.ReplaceIfNotNil(dto.Lastname, &entity.Lastname)
	utils.ReplaceIfNotNil(dto.Patronymic, &entity.Patronymic)
	utils.ReplaceIfNotNil(dto.Roles, &entity.Roles)

	return entity
}

package usecases

import (
	"github.com/jemilsonluis/internals/modules/users/domain/dto"
	"github.com/jemilsonluis/internals/modules/users/domain/entity"
)

type ICreateUserUseCase interface {
	Execute(params dto.CreateUserDTO) (*entity.UserEntity, error)
}

package usecases

import (
	"github.com/jemilsonluis/internals/modules/users/domain/dto"
	"github.com/jemilsonluis/internals/modules/users/domain/entity"
)

type IUpdateUserUseCase interface {
	Execute(params dto.UpdateUserDTO) (*entity.UserEntity, error)
}

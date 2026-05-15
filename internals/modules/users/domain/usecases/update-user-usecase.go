package usecases

import (
	"github.com/jemilsonluis/internals/_shared/params"
	"github.com/jemilsonluis/internals/modules/users/domain/entity"
)

type IUpdateUserUseCase interface {
	Execute(params params.UpdateUserParams) (*entity.UserEntity, error)
}

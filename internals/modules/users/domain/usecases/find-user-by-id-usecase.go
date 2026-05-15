package usecases

import (
	"github.com/jemilsonluis/internals/modules/users/domain/entity"
)

type IFindUserByIdUseCase interface {
	Execute(userId string) (*entity.UserEntity, error)
}

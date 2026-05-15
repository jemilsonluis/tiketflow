package usecases

import (
	"github.com/jemilsonluis/internals/modules/users/domain/entity"
)

type IFindUserByEmailUseCase interface {
	Execute(email string) (*entity.UserEntity, error)
}

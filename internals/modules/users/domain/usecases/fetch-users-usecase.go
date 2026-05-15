package usecases

import (
	"github.com/jemilsonluis/internals/modules/users/domain/entity"
)

type IFetchUsersUseCase interface {
	Execute() ([]*entity.UserEntity, error)
}

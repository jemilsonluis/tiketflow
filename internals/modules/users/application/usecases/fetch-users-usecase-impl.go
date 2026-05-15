package usecases

import (
	"github.com/jemilsonluis/internals/modules/users/application/repository"
	"github.com/jemilsonluis/internals/modules/users/domain/entity"
)

type FetchUsersUseCaseImpl struct {
	Repo repository.IUserRepository
}

func (uc *FetchUsersUseCaseImpl) Execute() ([]*entity.UserEntity, error) {
	return uc.Repo.FetchUsers()
}

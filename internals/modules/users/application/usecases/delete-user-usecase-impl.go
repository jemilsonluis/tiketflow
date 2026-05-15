package usecases

import (
	"github.com/jemilsonluis/internals/modules/users/application/repository"
)

type DeleteUserUseCaseImpl struct {
	Repo repository.IUserRepository
}

func (uc *DeleteUserUseCaseImpl) Execute(userId string) error {
	return uc.Repo.Delete(userId)
}

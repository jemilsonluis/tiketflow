package usecases

import (
	"github.com/jemilsonluis/internals/modules/users/application/repository"
	"github.com/jemilsonluis/internals/modules/users/domain/entity"
)

type FindUserByIdUseCaseImpl struct {
	Repo repository.IUserRepository
}

func (uc *FindUserByIdUseCaseImpl) Execute(userId string) (*entity.UserEntity, error) {
	return uc.Repo.FindUserById(userId)
}

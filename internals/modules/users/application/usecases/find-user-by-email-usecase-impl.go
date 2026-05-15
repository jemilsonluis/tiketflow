package usecases

import (
	"github.com/jemilsonluis/internals/modules/users/application/repository"
	"github.com/jemilsonluis/internals/modules/users/domain/entity"
)

type FindUserByEmailUseCaseImpl struct {
	Repo repository.IUserRepository
}

func (uc *FindUserByEmailUseCaseImpl) Execute(email string) (*entity.UserEntity, error) {
	return uc.Repo.FindUserByEmail(email)
}

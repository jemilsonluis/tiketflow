package usecases

import (
	"github.com/jemilsonluis/internals/_shared/params"
	"github.com/jemilsonluis/internals/modules/users/application/repository"
	"github.com/jemilsonluis/internals/modules/users/domain/entity"
)

type UpdateUserUseCaseImpl struct {
	Repo repository.IUserRepository
}

func (uc *UpdateUserUseCaseImpl) Execute(params params.UpdateUserParams) (*entity.UserEntity, error) {
	return uc.Repo.Update(params)
}

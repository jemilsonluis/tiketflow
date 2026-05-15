package usecases

import (
	"github.com/jemilsonluis/internals/modules/users/application/repository"
	"github.com/jemilsonluis/internals/modules/users/domain/dto"
	"github.com/jemilsonluis/internals/modules/users/domain/entity"
)

type CreateUserUseCaseImpl struct {
	Repo repository.IUserRepository
}

func (uc *CreateUserUseCaseImpl) Execute(params dto.CreateUserDTO) (*entity.UserEntity, error) {
	return uc.Repo.Create(params)
}

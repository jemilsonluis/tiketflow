package usecases

import (
	"github.com/jemilsonluis/internals/modules/users/application/repository"
	"github.com/jemilsonluis/internals/modules/users/domain/dto"
	"github.com/jemilsonluis/internals/modules/users/domain/entity"
)

type UpdateUserUseCaseImpl struct {
	Repo repository.IUserRepository
}

func (uc *UpdateUserUseCaseImpl) Execute(params dto.UpdateUserDTO) (*entity.UserEntity, error) {
	return uc.Repo.Update(params)
}

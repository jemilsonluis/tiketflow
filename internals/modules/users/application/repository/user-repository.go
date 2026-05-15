package repository

import (
	"github.com/jemilsonluis/internals/_shared/params"
	"github.com/jemilsonluis/internals/modules/users/domain/dto"
	"github.com/jemilsonluis/internals/modules/users/domain/entity"
)

type IUserRepository interface {
	Create(params dto.CreateUserDTO) (*entity.UserEntity, error)
	FetchUsers() ([]*entity.UserEntity, error)
	FindUserById(userId string) (*entity.UserEntity, error)
	FindUserByEmail(email string) (*entity.UserEntity, error)
	Update(params params.UpdateUserParams) (*entity.UserEntity, error)
	Delete(userId string) error
}

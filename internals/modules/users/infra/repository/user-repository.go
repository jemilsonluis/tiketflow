package repository

import (
	"github.com/jemilsonluis/internals/modules/users/application/repository"
	"github.com/jemilsonluis/internals/modules/users/domain/dto"
	"github.com/jemilsonluis/internals/modules/users/domain/entity"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() repository.IUserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Create(params dto.CreateUserDTO) (*entity.UserEntity, error) {
	return nil, nil
}

func (u *UserRepositoryImpl) FetchUsers() ([]*entity.UserEntity, error) {
	return nil, nil
}

func (u *UserRepositoryImpl) FindUserById(userId string) (*entity.UserEntity, error) {
	return nil, nil
}

func (u *UserRepositoryImpl) FindUserByEmail(email string) (*entity.UserEntity, error) {
	return nil, nil
}

func (u *UserRepositoryImpl) Update(params dto.UpdateUserDTO) (*entity.UserEntity, error) {
	return nil, nil
}

func (u *UserRepositoryImpl) Delete(userId string) error {
	return nil
}

package factory

import (
	"github.com/jemilsonluis/internals/modules/users/application/repository"
	"github.com/jemilsonluis/internals/modules/users/application/usecases"
	iusecase "github.com/jemilsonluis/internals/modules/users/domain/usecases"
)

type IUserFactory interface {
	CreateUserUsecase() iusecase.ICreateUserUseCase
	FetchUsersUsecase() iusecase.IFetchUsersUseCase
	FindUserUsecase() iusecase.IFindUserByIdUseCase
	FindUserByEmailUsecase() iusecase.IFindUserByEmailUseCase
	UpdateUserUsecase() iusecase.IUpdateUserUseCase
	DeleteUserUsecase() iusecase.IDeleteUserUseCase
}

type UserFactoryImpl struct {
	repo repository.IUserRepository
}

func NewUserFactoryImpl(repo repository.IUserRepository) IUserFactory {
	return &UserFactoryImpl{repo: repo}
}

func (u *UserFactoryImpl) CreateUserUsecase() iusecase.ICreateUserUseCase {
	return &usecases.CreateUserUseCaseImpl{Repo: u.repo}
}

func (u *UserFactoryImpl) FetchUsersUsecase() iusecase.IFetchUsersUseCase {
	return &usecases.FetchUsersUseCaseImpl{Repo: u.repo}
}

func (u *UserFactoryImpl) FindUserUsecase() iusecase.IFindUserByIdUseCase {
	return &usecases.FindUserByIdUseCaseImpl{Repo: u.repo}
}

func (u *UserFactoryImpl) FindUserByEmailUsecase() iusecase.IFindUserByEmailUseCase {
	return &usecases.FindUserByEmailUseCaseImpl{Repo: u.repo}
}

func (u *UserFactoryImpl) UpdateUserUsecase() iusecase.IUpdateUserUseCase {
	return &usecases.UpdateUserUseCaseImpl{Repo: u.repo}
}

func (u *UserFactoryImpl) DeleteUserUsecase() iusecase.IDeleteUserUseCase {
	return &usecases.DeleteUserUseCaseImpl{Repo: u.repo}
}

package handlers

import (
	"net/http"

	"github.com/jemilsonluis/internals/modules/users/domain/dto"
	enum "github.com/jemilsonluis/internals/modules/users/domain/enums"
	"github.com/jemilsonluis/internals/modules/users/factory"
)

type IUserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Fetch(w http.ResponseWriter, r *http.Request)
	FetchByEmail(w http.ResponseWriter, r *http.Request)
	Find(w http.ResponseWriter, r *http.Request)
}

type UserHandlerImpl struct {
	factory factory.IUserFactory
}

func NewUserhandlerImpl(factory factory.IUserFactory) IUserHandler {
	return &UserHandlerImpl{
		factory: factory,
	}
}

func (h *UserHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
	usecase := h.factory.CreateUserUsecase()
	params := dto.CreateUserDTO{}
	usecase.Execute(params)
}
func (h *UserHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	usecase := h.factory.DeleteUserUsecase()
	usecase.Execute("")
}
func (h *UserHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	usecase := h.factory.UpdateUserUsecase()
	params := dto.UpdateUserDTO{
		UserId: "",
		Name:   "",
		Role:   enum.AdminRoleEnum,
	}

	usecase.Execute(params)
}
func (h *UserHandlerImpl) Fetch(w http.ResponseWriter, r *http.Request) {
	usecase := h.factory.FetchUsersUsecase()
	usecase.Execute()
}
func (h *UserHandlerImpl) FetchByEmail(w http.ResponseWriter, r *http.Request) {
	usecase := h.factory.FindUserByEmailUsecase()
	usecase.Execute("")
}
func (h *UserHandlerImpl) Find(w http.ResponseWriter, r *http.Request) {
	usecase := h.factory.FindUserUsecase()
	usecase.Execute("")
}

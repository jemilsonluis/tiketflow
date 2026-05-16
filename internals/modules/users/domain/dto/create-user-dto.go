package dto

import enum "github.com/jemilsonluis/internals/modules/users/domain/enums"


type CreateUserDTO struct {
	Name  string
	Email string
	Role  enum.UserRoleEnum
}

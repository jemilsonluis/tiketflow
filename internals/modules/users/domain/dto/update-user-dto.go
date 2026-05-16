package dto

import enum "github.com/jemilsonluis/internals/modules/users/domain/enums"


type UpdateUserDTO struct {
	UserId string
	Name   string
	Role   enum.UserRoleEnum
}

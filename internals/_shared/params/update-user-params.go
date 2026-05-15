package params

import "github.com/jemilsonluis/internals/_shared/enum"

type UpdateUserParams struct {
	UserId string
	Name   string
	Role   enum.UserRoleEnum
}

package enum

type UserRoleEnum string

const (
	AdminRoleEnum     UserRoleEnum = "ADMIN"
	OrganizerRoleEnum UserRoleEnum = "ORGANIZER"
	CustomerRoleEnum  UserRoleEnum = "CUSTOMER"
)

func IsUserRoleEnum(role string) bool {
	switch role {
	case
		string(AdminRoleEnum),
		string(OrganizerRoleEnum),
		string(CustomerRoleEnum):
		return true
	default:
		return false
	}
}

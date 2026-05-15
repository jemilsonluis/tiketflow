package entity

type UserEntity struct {
	Id    string
	Name  string
	Email string
	Role  string
}

func NewUserEntity() (*UserEntity, error) {

	return &UserEntity{}, nil
}

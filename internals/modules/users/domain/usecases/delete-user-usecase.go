package usecases

type IDeleteUserUseCase interface {
	Execute(userId string) error
}

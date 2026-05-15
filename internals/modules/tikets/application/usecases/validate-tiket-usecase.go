package usecases

import "github.com/jemilsonluis/internals/modules/tikets/application/repository"

type ValidateTiketUseCaseImpl struct {
	Repo repository.ITiketRepository
}

func (v *ValidateTiketUseCaseImpl) Exec(tiketId string) (bool, error) {
	return v.Repo.ValidateTiket(tiketId)
}

package usecases

import (
	"github.com/jemilsonluis/internals/modules/tikets/application/repository"
	"github.com/jemilsonluis/internals/modules/tikets/domain/entity"
)

type FetchTiketByIdUseCaseImpl struct {
	Repo repository.ITiketRepository
}

func (f *FetchTiketByIdUseCaseImpl) Exec(id string) (*entity.TiketEntity, error) {
	return f.Repo.FindTiketById(id)
}

package usecases

import (
	"github.com/jemilsonluis/internals/modules/tikets/application/repository"
	"github.com/jemilsonluis/internals/modules/tikets/domain/entity"
)

type FetchTiketByCodeUseCaseImpl struct {
	repo repository.ITiketRepository
}

func (f *FetchTiketByCodeUseCaseImpl) Exec(code string) (*entity.TiketEntity, error) {
	return f.repo.FindTiketByCode(code)
}

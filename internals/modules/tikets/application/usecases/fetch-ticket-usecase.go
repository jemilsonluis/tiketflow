package usecases

import (
	"github.com/jemilsonluis/internals/modules/tikets/application/repository"
	"github.com/jemilsonluis/internals/modules/tikets/domain/entity"
	"github.com/jemilsonluis/internals/modules/tikets/domain/enum"
)

type FetchTiketsUseCaseImpl struct {
	Repo repository.ITiketRepository
}

func (f *FetchTiketsUseCaseImpl) Exec(tiketType enum.TiketTypeEnum) ([]*entity.TiketEntity, error) {
	// caso tenha o tiketType, busque por ele
	return f.Repo.FetchTikets(tiketType)
}

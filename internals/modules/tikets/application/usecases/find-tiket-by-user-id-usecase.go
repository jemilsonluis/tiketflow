package usecases

import (
	"github.com/jemilsonluis/internals/modules/tikets/application/repository"
	"github.com/jemilsonluis/internals/modules/tikets/domain/entity"
)

type FindTiketByUserIdUseCaseImpl struct {
	Repo repository.ITiketRepository
}

func (f *FindTiketByUserIdUseCaseImpl) Exec(userId string) (*entity.TiketEntity, error) {
	return f.Repo.FindTiketByUserId(userId)
}

package usecases

import (
	"github.com/jemilsonluis/internals/_shared/params"
	"github.com/jemilsonluis/internals/modules/tikets/domain/entity"
)

type ICreateTiketUseCase interface {
	Exec(params params.CreateTiketParams) (entity.TiketEntity, error)
}

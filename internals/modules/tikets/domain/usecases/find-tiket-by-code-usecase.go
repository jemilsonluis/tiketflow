package usecases

import "github.com/jemilsonluis/internals/modules/tikets/domain/entity"

type IFindTiketByCodeUseCase interface {
	Exec(code string) (entity.TiketEntity, error)
}
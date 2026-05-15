package usecases

import "github.com/jemilsonluis/internals/modules/tikets/domain/entity"

type IFindTiketByIdUseCase interface {
	Exec(tiketId string) (entity.TiketEntity, error)
}

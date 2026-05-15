package usecases

import "github.com/jemilsonluis/internals/modules/tikets/domain/entity"

type IFindTiketByUserIdUseCase interface {
	Exec(userId string) ([]entity.TiketEntity, error)
}

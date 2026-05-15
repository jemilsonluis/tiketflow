package usecases

import "github.com/jemilsonluis/internals/modules/tikets/domain/entity"

type IValidateTikettUseCase interface {
	Exec(tiketId string) (entity.TiketEntity, error)
}

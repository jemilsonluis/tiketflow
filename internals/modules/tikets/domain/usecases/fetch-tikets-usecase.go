package usecases

import (
	"github.com/jemilsonluis/internals/modules/tikets/domain/entity"
	"github.com/jemilsonluis/internals/modules/tikets/domain/enum"
)

type IFetchTiketsUseCase interface {
	Exec(tiketType enum.TiketTypeEnum) ([]entity.TiketEntity, error)
}

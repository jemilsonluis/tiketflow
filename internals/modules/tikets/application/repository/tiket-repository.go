package repository

import (
	"github.com/jemilsonluis/internals/modules/tikets/domain/dto"
	"github.com/jemilsonluis/internals/modules/tikets/domain/entity"
	"github.com/jemilsonluis/internals/modules/tikets/domain/enum"
)

type ITiketRepository interface {
	Create(params dto.CreateTiketDTO) (*entity.TiketEntity, error)
	FetchTikets(tiketType enum.TiketTypeEnum) ([]*entity.TiketEntity, error)
	FindTiketByCode(code string) (*entity.TiketEntity, error)
	FindTiketById(tiketId string) (*entity.TiketEntity, error)
	FindTiketByUserId(userId string) (*entity.TiketEntity, error)
	ValidateTiket(tiketId string) (bool, error)
}

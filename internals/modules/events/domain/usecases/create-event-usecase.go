package usecases

import (
	"github.com/jemilsonluis/internals/modules/events/domain/dto"
	"github.com/jemilsonluis/internals/modules/events/domain/entity"
)

type ICreateEventUseCase interface {
	Exec(params dto.CreateEventDTO) (entity.EventEntity, error)
}

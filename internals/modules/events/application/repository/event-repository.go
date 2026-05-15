package repository

import (
	"github.com/jemilsonluis/internals/_shared/params"
	"github.com/jemilsonluis/internals/modules/events/domain/dto"
	"github.com/jemilsonluis/internals/modules/events/domain/entity"
)

type IEventRepository interface {
	Create(params dto.CreateEventDTO) (entity.EventEntity, error)
	FetchEvents(params params.FetchEventsParams) ([]entity.EventEntity, error)
	FindEvent(eventId string) (entity.EventEntity, error)
}

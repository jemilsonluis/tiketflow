package usecases

import "github.com/jemilsonluis/internals/modules/events/domain/entity"

type IFindEventUseCase interface {
	Exec(eventId string) (entity.EventEntity, error)
}

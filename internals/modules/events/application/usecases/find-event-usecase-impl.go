package usecases

import (
	eventRepo "github.com/jemilsonluis/internals/modules/events/application/repository"
	"github.com/jemilsonluis/internals/modules/events/domain/entity"
)

type FindEventUseCaseImpl struct {
	repo eventRepo.IEventRepository
}

func (uc *FetchEventsUseCaseImpl) FindEvent(eventId string) (entity.EventEntity, error) {
	return uc.repo.FindEvent(eventId)
}
	
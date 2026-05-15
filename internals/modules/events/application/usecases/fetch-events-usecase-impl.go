package usecases

import (
	"github.com/jemilsonluis/internals/_shared/params"
	eventRepo "github.com/jemilsonluis/internals/modules/events/application/repository"
	"github.com/jemilsonluis/internals/modules/events/domain/entity"
)

type FetchEventsUseCaseImpl struct {
	repo eventRepo.IEventRepository
}

func (uc *FetchEventsUseCaseImpl) FetchEvents(params params.FetchEventsParams) ([]entity.EventEntity, error) {
	return uc.repo.FetchEvents(params)
}

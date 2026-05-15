package usecases

import (
	eventRepo "github.com/jemilsonluis/internals/modules/events/application/repository"
	"github.com/jemilsonluis/internals/modules/events/domain/dto"
	"github.com/jemilsonluis/internals/modules/events/domain/entity"
)

type CreateEventUseCaseImpl struct {
	Repo eventRepo.IEventRepository
}

func (uc *CreateEventUseCaseImpl) Create(params dto.CreateEventDTO) (entity.EventEntity, error) {
	return uc.Repo.Create(params)
}

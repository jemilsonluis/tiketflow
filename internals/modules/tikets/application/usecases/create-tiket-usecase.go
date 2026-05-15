package usecases

import (
	"errors"

	"github.com/jemilsonluis/internals/_shared/enum"
	eventRepo "github.com/jemilsonluis/internals/modules/events/application/repository"
	"github.com/jemilsonluis/internals/modules/tikets/application/handler"
	"github.com/jemilsonluis/internals/modules/tikets/application/repository"
	"github.com/jemilsonluis/internals/modules/tikets/domain/dto"
	"github.com/jemilsonluis/internals/modules/tikets/domain/entity"
)

type CreateTiketUseCaseImpl struct {
	Repo      repository.ITiketRepository
	EventRepo eventRepo.IEventRepository
}

func (c *CreateTiketUseCaseImpl) Exec(params dto.CreateTiketDTO) (*entity.TiketEntity, error) {
	event, err := c.EventRepo.FindEvent(params.EventId) // Validate if the event exists
	if err != nil {
		return nil, err
	}

	// Validate if the event code already exists
	for i := 0; i < 5; i++ { // try to generate a unique code up to 5 times
		eventCode := handler.GenerateEventCode(event.Title)

		tiket, err := c.Repo.FindTiketByCode(eventCode)
		if err != nil {
			return nil, err
		}

		if tiket == nil {
			params.Code = eventCode
			return c.Repo.Create(params)
		}
	}

	return nil, errors.New(string(enum.GenerateCodeFailed))
}

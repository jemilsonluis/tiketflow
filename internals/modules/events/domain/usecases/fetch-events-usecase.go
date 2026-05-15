package usecases

import (
	"github.com/jemilsonluis/internals/_shared/params"
	"github.com/jemilsonluis/internals/modules/events/domain/entity"
)

type IFetchEventsUseCase interface {
	Exec(params params.FetchEventsParams) ([]entity.EventEntity, error)
}

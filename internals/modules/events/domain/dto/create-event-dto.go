package dto

import (
	"time"

	"github.com/jemilsonluis/internals/_shared/types"
	"github.com/jemilsonluis/internals/modules/events/domain/enum"
)

type CreateEventDTO struct {
	Title       string
	Description string
	Banner      string
	Capacity    int
	Status      enum.EventStatusEnum
	Location    types.LocationType
	StartsAt    time.Time
	EndsAt      time.Time
	EventTypeId string
	OrganizerId string
}

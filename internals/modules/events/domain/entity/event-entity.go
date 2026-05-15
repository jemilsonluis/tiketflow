package entity

import (
	"errors"
	"time"

	"github.com/jemilsonluis/internals/_shared/identity"
	"github.com/jemilsonluis/internals/_shared/types"
	"github.com/jemilsonluis/internals/modules/events/domain/dto"
	"github.com/jemilsonluis/internals/modules/events/domain/enum"
)

type EventEntity struct {
	Id          string
	Title       string
	Description string
	Banner      string
	Capacity    int
	Location    types.LocationType
	Status      enum.EventStatusEnum
	StartsAt    time.Time
	EndsAt      time.Time
	EventTypeId string
	OrganizerId string
}

func NewEventEntity(e dto.CreateEventDTO) (*EventEntity, error) {
	if e.Title == "" {
		return nil, errors.New(string(enum.InvalidTitle))
	}
	if e.Description == "" {
		return nil, errors.New(string(enum.InvalidDescription))
	}
	if !types.IsValidLocationType(e.Location) {
		return nil, errors.New(string(enum.InvalidLocation))
	}
	if e.Capacity <= 0 {
		return nil, errors.New(string(enum.InvalidCapacity))
	}
	if !enum.IsEventStatusEnum(string(e.Status)) {
		return nil, errors.New(string(enum.InvalidStatus))
	}
	if e.StartsAt.IsZero() {
		return nil, errors.New(string(enum.InvalidStartsAt))
	}
	if e.EndsAt.IsZero() {
		return nil, errors.New(string(enum.InvalidEndsAt))
	}
	if e.StartsAt.After(e.EndsAt) {
		return nil, errors.New(string(enum.InvalidStartsAt))
	}
	if e.EventTypeId == "" {
		return nil, errors.New(string(enum.InvalidEventTypeId))
	}
	if e.OrganizerId == "" {
		return nil, errors.New(string(enum.InvalidOrganizerId))
	}

	return &EventEntity{
		Id:          identity.GenerateId(),
		Title:       e.Title,
		Description: e.Description,
		Location:    e.Location,
		Banner:      e.Banner,
		Capacity:    e.Capacity,
		Status:      enum.EventStatusEnum(e.Status),
		StartsAt:    e.StartsAt,
		EndsAt:      e.EndsAt,
		OrganizerId: e.OrganizerId,
	}, nil
}

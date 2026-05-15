package entity

import (
	"errors"
	"time"

	"github.com/jemilsonluis/internals/_shared/identity"
	"github.com/jemilsonluis/internals/modules/tikets/domain/dto"
	"github.com/jemilsonluis/internals/modules/tikets/domain/enum"
)

type TiketEntity struct {
	Id        string
	Code      string
	UserId    string
	EventId   string
	TiketType enum.TiketTypeEnum
	ExpiresAt time.Time
}

func NewTiketEntity(e dto.CreateTiketDTO) (*TiketEntity, error) {

	if e.Code == "" {
		return nil, errors.New(string(enum.InvalidCode))
	}
	if e.UserId == "" {
		return nil, errors.New(string(enum.InvalidUserId))
	}
	if e.EventId == "" {
		return nil, errors.New(string(enum.InvalidEventId))
	}
	if !enum.IsTiketTypeEnum(string(e.TiketType)) {
		return nil, errors.New(string(enum.InvalidTiketType))
	}
	if e.ExpiresAt.IsZero() {
		return nil, errors.New(string(enum.InvalidExpiresAt))
	}
	if e.ExpiresAt.Before(time.Now()) {
		return nil, errors.New(string(enum.InvalidExpiresAt))
	}

	return &TiketEntity{
		Id:        identity.GenerateId(),
		Code:      e.Code,
		UserId:    e.UserId,
		EventId:   e.EventId,
		TiketType: e.TiketType,
		ExpiresAt: e.ExpiresAt,
	}, nil
}

package dto

import (
	"time"

	"github.com/jemilsonluis/internals/modules/tikets/domain/enum"
)

type CreateTiketDTO struct {
	Code      string
	UserId    string
	EventId   string
	TiketType enum.TiketTypeEnum
	ExpiresAt time.Time
}

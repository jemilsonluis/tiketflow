package params

import (
	"time"

)

type CreateTiketParams struct {
	Code string
	UserId string
	EventId string
	TiketType enum .TiketTypeEnum
	ExpiresAt time.Time
}
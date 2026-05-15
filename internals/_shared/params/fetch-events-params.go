package params

import (
	"time"

	"github.com/jemilsonluis/internals/_shared/types"
)

type FetchEventsParams struct {
	Date      time.Time
	Location  types.LocationType
	EventType string
	Price     int
}

package enum

type EventErrorsEnum string

const (
	InvalidTitle       EventErrorsEnum = "ERR_INVALID_TITLE"
	InvalidDescription EventErrorsEnum = "ERR_INVALID_DESCRIPTION"
	InvalidLocation    EventErrorsEnum = "ERR_INVALID_LOCATION"
	InvalidBanner      EventErrorsEnum = "ERR_INVALID_BANNER"
	InvalidCapacity    EventErrorsEnum = "ERR_INVALID_CAPACITY"
	InvalidStatus      EventErrorsEnum = "ERR_INVALID_STATUS"
	InvalidStartsAt    EventErrorsEnum = "ERR_INVALID_STARTS_AT"
	InvalidEndsAt      EventErrorsEnum = "ERR_INVALID_ENDS_AT"
	InvalidEventTypeId EventErrorsEnum = "ERR_INVALID_EVENT_TYPE_ID"
	InvalidOrganizerId EventErrorsEnum = "ERR_INVALID_ORGANIZER_ID"
)

func IsEventErrorsEnum(err string) bool {
	switch err {
	case
		string(InvalidTitle),
		string(InvalidDescription),
		string(InvalidLocation),
		string(InvalidBanner),
		string(InvalidCapacity),
		string(InvalidStatus),
		string(InvalidStartsAt),
		string(InvalidEndsAt),
		string(InvalidEventTypeId),
		string(InvalidOrganizerId):
		return true
	default:
		return false
	}
}

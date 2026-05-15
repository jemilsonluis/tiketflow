package enum

type ErrorsEnum string

const (
	InvalidTitle ErrorsEnum = "ERR_INVALID_TITLE"
	InvalidDescription ErrorsEnum = "ERR_INVALID_DESCRIPTION"
	InvalidLocation ErrorsEnum = "ERR_INVALID_LOCATION"
	InvalidBanner ErrorsEnum = "ERR_INVALID_BANNER"
	InvalidCapacity ErrorsEnum = "ERR_INVALID_CAPACITY"
	InvalidStatus ErrorsEnum = "ERR_INVALID_STATUS"
	InvalidStartsAt ErrorsEnum = "ERR_INVALID_STARTS_AT"
	InvalidEndsAt ErrorsEnum = "ERR_INVALID_ENDS_AT"
	InvalidEventTypeId ErrorsEnum = "ERR_INVALID_EVENT_TYPE_ID"
	InvalidOrganizerId ErrorsEnum = "ERR_INVALID_ORGANIZER_ID"

	InvalidCode ErrorsEnum = "ERR_INVALID_CODE"
	InvalidUserId ErrorsEnum = "ERR_INVALID_USER_ID"
	InvalidEventId ErrorsEnum = "ERR_INVALID_EVENT_ID"
	InvalidTiketType ErrorsEnum = "ERR_INVALID_TIKET_TYPE"
	InvalidExpiresAt ErrorsEnum = "ERR_INVALID_EXPIRES_AT"

	GenerateCodeFailed ErrorsEnum = "ERR_GENERATE_CODE_FAILED"
)

func IsErrorEnum(err string) bool {
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
		string(InvalidOrganizerId),
		string(InvalidCode),
		string(InvalidUserId),
		string(InvalidEventId),
		string(InvalidTiketType),
		string(InvalidExpiresAt),
		string(GenerateCodeFailed):
		return true
	default:
		return false
	}
}

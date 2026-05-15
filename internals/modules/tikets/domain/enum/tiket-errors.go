package enum

type ErrorsEnum string

const (
	InvalidCode      ErrorsEnum = "ERR_INVALID_CODE"
	InvalidUserId    ErrorsEnum = "ERR_INVALID_USER_ID"
	InvalidEventId   ErrorsEnum = "ERR_INVALID_EVENT_ID"
	InvalidTiketType ErrorsEnum = "ERR_INVALID_TIKET_TYPE"
	InvalidExpiresAt ErrorsEnum = "ERR_INVALID_EXPIRES_AT"
)

func IsErrorEnum(err string) bool {
	switch err {
	case
		string(InvalidCode),
		string(InvalidUserId),
		string(InvalidEventId),
		string(InvalidTiketType),
		string(InvalidExpiresAt):
		return true
	default:
		return false
	}
}

package enum

type EventStatusEnum string

const (
	EventStatusDraft     EventStatusEnum = "DRAFT"
	EventStatusPublished EventStatusEnum = "PUBLISHED"
	EventStatusCanceled  EventStatusEnum = "CANCELED"
	EventStatusFinished  EventStatusEnum = "FINISHED"
)

func IsEventStatusEnum(status string) bool {
	switch status {
	case
		string(EventStatusCanceled),
		string(EventStatusDraft),
		string(EventStatusPublished),
		string(EventStatusFinished):
		return true
	default:
		return false
	}
}

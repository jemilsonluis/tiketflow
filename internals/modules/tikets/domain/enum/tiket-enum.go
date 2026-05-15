package enum

type TiketTypeEnum string

const (
	TiketTypeRegular TiketTypeEnum = "REGULAR"
	TiketTypeStudent TiketTypeEnum = "STUDENT"
	TiketTypeVIP     TiketTypeEnum = "VIP"
)

func IsTiketTypeEnum(tiketType string) bool {
	switch tiketType {
	case
		string(TiketTypeRegular),
		string(TiketTypeStudent),
		string(TiketTypeVIP):
		return true
	default:
		return false
	}
}

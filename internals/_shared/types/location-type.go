package types

type LocationType struct {
	Country     string
	Province    string
	Municipal   string
	Adress      string
	MoreDetails string
}

func IsValidLocationType(l LocationType) bool {
	if l.Country == "" {
		return false
	}

	if l.Province == "" {
		return false
	}

	if l.Municipal == "" {
		return false
	}

	if l.Adress == "" {
		return false
	}

	if l.MoreDetails == "" {
		return false
	}

	return true
}

package aoe

import "time"

type defaultValidator struct {
}

func NewDefaultValidator() *defaultValidator {
	return &defaultValidator{}
}

// Valid will check if given record complies with AOE data rules
func (val *defaultValidator) Valid(record *Record) bool {
	releaseDate, _ := time.Parse("2006.01.02", record.ReleaseDate)
	limitDate, _ := time.Parse("2006.01.02", "2001.01.01")
	return len(record.TrackListing.Tracks) >= 10 && releaseDate.Before(limitDate)
}

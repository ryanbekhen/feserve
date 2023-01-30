package timeutils

import "time"

func Location(timezone string) *time.Location {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		loc, _ = time.LoadLocation("UTC")
	}
	return loc
}

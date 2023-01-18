package logger

import "time"

func loadLocation(timezone string) *time.Location {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		loc, _ = time.LoadLocation("UTC")
	}
	return loc
}

package timeutils

import (
	"math"
	"time"
)

type Duration struct {
	different time.Duration
}

func DiffCurtime(datetime time.Time, location *time.Location) *Duration {
	currentTime := time.Now().In(location)
	diff := datetime.Sub(currentTime)
	return &Duration{diff}
}

func (diff *Duration) Days() float64 {
	days := diff.different.Hours() / 24
	return math.Abs(days)
}

func (diff *Duration) Hours() float64 {
	return math.Abs(diff.different.Hours())
}

func (diff *Duration) Minutes() float64 {
	return math.Abs(diff.different.Minutes())
}

func (diff *Duration) Seconds() float64 {
	return math.Abs(diff.different.Seconds())
}

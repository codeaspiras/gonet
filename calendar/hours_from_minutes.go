package calendar

import "math"

func HoursFromMinutes(minutes float64) float64 {
	return math.Floor(minutes / 60)
}

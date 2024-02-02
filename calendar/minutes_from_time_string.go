package calendar

import "time"

func MinutesFromTimeString(t string) int {
	timestamp, err := time.Parse(TimeLayout, t)
	if err != nil {
		panic(err)
	}

	return (timestamp.Hour() * 60) + timestamp.Minute()
}

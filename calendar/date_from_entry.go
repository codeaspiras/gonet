package calendar

import (
	"strconv"
	"time"
)

func WeekdayFromEntry(dateString string) string {
	date, err := time.Parse(YMDLayout, dateString)
	if err != nil {
		panic(err)
	}

	return strconv.Itoa(int(date.Weekday()))
}

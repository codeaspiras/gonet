package main

import (
	"gonet/calendar"
	"gonet/utils"
)

func GetWorkingMinutesPerWeekday() map[string][][]int {
	var workingHoursPerWeekday map[string][][]string
	workingMinutesPerWeekday := make(map[string][][]int)
	utils.GetJSON("horas-uteis.json", &workingHoursPerWeekday)

	for weekday, workingHours := range workingHoursPerWeekday {
		workingMinutesPerWeekday[weekday] = make([][]int, len(workingHours))
		for rangeIndex, hoursRange := range workingHours {
			start := calendar.MinutesFromTimeString(hoursRange[0])
			end := calendar.MinutesFromTimeString(hoursRange[1])
			if start > end {
				// switch start with end
				start -= end
				end += start
				start = end - start
			}
			workingMinutesPerWeekday[weekday][rangeIndex] = []int{start, end}
		}
	}

	return workingMinutesPerWeekday
}

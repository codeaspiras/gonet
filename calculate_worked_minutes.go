package main

import (
	"gonet/calendar"
	"gonet/utils"
)

type Result struct {
	WorkingMinutes float64
	ExtraMinutes   float64
}

func CalculateWorkedMinutes(
	workingMinutesPerWeekday map[string][][]int,
) *Result {
	var timeEntriesPerDate map[string][][]string
	utils.GetJSON("registros-de-ponto.json", &timeEntriesPerDate)

	result := Result{
		WorkingMinutes: 0,
		ExtraMinutes:   0,
	}
	var expectedWorkingMinutes float64 = 0
	for dateString, entries := range timeEntriesPerDate {
		weekday := calendar.WeekdayFromEntry(dateString)
		workingMinuteRanges, hasMinuteRanges := workingMinutesPerWeekday[weekday]
		if hasMinuteRanges {
			for _, workingMinuteRange := range workingMinuteRanges {
				expectedWorkingMinutes += float64(workingMinuteRange[1] - workingMinuteRange[0])
			}
		}

		for _, entry := range entries {
			start := calendar.MinutesFromTimeString(entry[0])
			end := calendar.MinutesFromTimeString(entry[1])
			totalDuration := float64(end - start)

			if !hasMinuteRanges || len(workingMinuteRanges) == 0 {
				result.ExtraMinutes += totalDuration
				continue
			}

			var innerWorkingMinutes float64 = 0
			var innerExtraMinutes float64 = 0
			for _, workingMinuteRange := range workingMinuteRanges {
				// current range happens before this working-range
				if start < workingMinuteRange[0] && end < workingMinuteRange[0] {
					continue // check next range
				}

				// current range happens after this working-range
				if start > workingMinuteRange[1] && end > workingMinuteRange[1] {
					continue // check next range
				}

				// current range happens between this working-range
				if start >= workingMinuteRange[0] && end <= workingMinuteRange[1] {
					innerWorkingMinutes += totalDuration
					break
				}

				// this working-range happens between current range
				if workingMinuteRange[0] >= start && workingMinuteRange[1] <= end {
					innerWorkingMinutes += float64(workingMinuteRange[1] - workingMinuteRange[0])
					break
				}

				// current range is around this working-range' start
				if start < workingMinuteRange[0] && end > workingMinuteRange[0] && end <= workingMinuteRange[1] {
					innerWorkingMinutes += float64(end - workingMinuteRange[0])
					continue // incomplete working minutes, check next ranges
				}

				// current range is around this working-range's end
				if start > workingMinuteRange[0] && start < workingMinuteRange[1] && end > workingMinuteRange[1] {
					innerWorkingMinutes += float64(workingMinuteRange[1] - start)
					continue // incomplete working minutes, check next ranges
				}

				// unexpected behavior
				panic("should not get here")
			}

			if innerWorkingMinutes < totalDuration {
				innerExtraMinutes += totalDuration - innerWorkingMinutes
			}

			result.WorkingMinutes += innerWorkingMinutes
			result.ExtraMinutes += innerExtraMinutes
		}
	}

	if result.WorkingMinutes < expectedWorkingMinutes && result.ExtraMinutes > 0 {
		debt := expectedWorkingMinutes - result.WorkingMinutes
		if debt <= result.ExtraMinutes {
			result.ExtraMinutes -= debt
			result.WorkingMinutes += debt
		} else {
			result.WorkingMinutes += result.ExtraMinutes
			result.ExtraMinutes = 0
		}
	}

	return &result
}

package main

import (
	"gonet/calendar"
	"gonet/utils"
)

func main() {
	utils.Echo("Bem-vindo(a) a calculadora de salário")
	name := askName()
	hourlyNet := askHourlyNet(name)

	minutelyNet := hourlyNet / 60

	utils.Echo("Lendo horas-uteis.json ...")
	workingMinutesPerWeekday := GetWorkingMinutesPerWeekday()

	utils.Echo("Lendo registros-de-ponto.json ...")
	result := CalculateWorkedMinutes(workingMinutesPerWeekday)

	utils.Echo("Horas úteis trabalhadas: %.0fh %dmin", calendar.HoursFromMinutes(result.WorkingMinutes), int(result.WorkingMinutes)%60)
	utils.Echo("Horas extras trabalhadas: %.0fh %dmin", calendar.HoursFromMinutes(result.ExtraMinutes), int(result.ExtraMinutes)%60)

	workingNet := result.WorkingMinutes * minutelyNet
	utils.Echo("Salário por horas úteis ... $ %f", workingNet)

	extraNet := result.ExtraMinutes * minutelyNet * 2
	utils.Echo("Salário por horas extras ... $ %f", extraNet)

	utils.Echo("Salário final: $ %f", workingNet+extraNet)
	utils.Finish()
}

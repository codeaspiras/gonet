package main

import (
	"gonet/utils"
	"strconv"
)

func askName() string {
	return utils.Ask("Qual o nome do(a) profissinoal?")
}

func askHourlyNet(name string) float64 {
	hourlyNetStr := utils.Ask("Qual o valor/hora de %s?", name)
	hourlyNet, err := strconv.ParseFloat(hourlyNetStr, 64)
	if err != nil {
		panic(err)
	}

	return hourlyNet
}

package main

import (
	"fmt"
)

func PrintFlightRow(FlightNum, FromCity, ToCity string, Time float64, RegistrationTable, Gate int, RegistrFinished bool) {
	if RegistrFinished {
		fmt.Printf(
			"| %s | %s—%s | регистрация закончилась, проходите к гейту: %d | длительность полёта: %.1f часа |\n",
			FlightNum,
			FromCity,
			ToCity,
			Gate,
			Time,
		)
	} else {
		fmt.Printf(
			"| %s | %s—%s | %d регистрация продолжается |\n",
			FlightNum,
			FromCity,
			ToCity,
			RegistrationTable,
		)
	}
}

func main() {
	var FlightNum, FromCity, ToCity string
	var Time float64
	var RegistrationTable, Gate int
	var RegistrFinished bool

	fmt.Scanln(&FlightNum)
	fmt.Scanln(&FromCity, &ToCity)
	fmt.Scanln(&Time)
	fmt.Scanln(&RegistrationTable)
	fmt.Scanln(&Gate)
	fmt.Scanln(&RegistrFinished)

	PrintFlightRow(FlightNum, FromCity, ToCity, Time, RegistrationTable, Gate, RegistrFinished)
}

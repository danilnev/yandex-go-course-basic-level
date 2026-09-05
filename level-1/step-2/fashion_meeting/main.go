package main

import (
	"fmt"
)

func main() {
	var symbol string
	var degree int

	fmt.Scan(&symbol)

	if symbol == "0" {
		fmt.Println("Стоит надеть куртку")
		return
	}

	fmt.Scanln(&degree)

	if symbol == "-" && degree > 5 {
		fmt.Println("Стоит надеть зимнюю куртку")
	} else if symbol == "-" || degree <= 9 {
		fmt.Println("Стоит надеть куртку")
	} else if degree <= 20 {
		fmt.Println("Стоит надеть штаны и кофту")
	} else {
		fmt.Println("Стоит надеть майку и шорты")
	}
}

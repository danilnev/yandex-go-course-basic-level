package main

import (
	"fmt"
)

func main() {
	var rubles, copecks, price int

	fmt.Scanln(&rubles, &copecks, &price)

	if (rubles + (copecks / 100)) >= price {
		fmt.Println("Сегодня будет вкусный кофе!")
		return
	}

	fmt.Println("Стоит подкопить")
}

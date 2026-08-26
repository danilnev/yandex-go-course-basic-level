package main

import (
	"fmt"
	"time"
)

func main() {
	var datetime string

	fmt.Scanln(&datetime)

	result_datetime, err := time.Parse("2006-01-02/15:04:05", datetime)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Текущее время %d часов, %d минут. Ты точно не забыл про важные дела на сегодня?\n", result_datetime.Hour(), result_datetime.Minute())
}

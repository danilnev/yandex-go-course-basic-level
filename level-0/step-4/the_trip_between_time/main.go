package main

import (
	"fmt"
	"time"
)

func main() {
	var str1, str2 string

	fmt.Scanln(&str1)
	fmt.Scanln(&str2)

	nowDate, nowErr := time.Parse("2006-01-02", str1)
	lateDate, lateErr := time.Parse("2006-01-02", str2)

	if nowErr != nil || lateErr != nil {
		fmt.Printf("Errors: %s\n\n%s\n", lateErr, nowErr)
		return
	}

	yearDifference := nowDate.Year() - lateDate.Year()

	fmt.Printf("%d year ago\n", yearDifference)
}

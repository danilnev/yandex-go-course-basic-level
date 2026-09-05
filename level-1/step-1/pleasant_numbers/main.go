package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Scanln(&num)

	if num == 0 {
		fmt.Println("Число 0")
	} else if -1 < num/10 && num/10 < 1 {
		fmt.Println("Число однозначное")
	} else if num%2 == 0 {
		fmt.Println("Число чётное")
	} else if num > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число красивое")
	}
}

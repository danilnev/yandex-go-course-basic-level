package main

import (
	"fmt"
)

func main() {
	var pass1, pass2 string

	fmt.Scanln(&pass1)
	fmt.Scanln(&pass2)

	if len(pass1) < 8 && len(pass2) < 8 {
		fmt.Println("Оба пароля ненадёжные")
	} else if len(pass1) < 8 {
		fmt.Println("Только второй пароль надёжный")
	} else if len(pass2) < 8 {
		fmt.Println("Только первый пароль надёжный")
	} else {
		fmt.Println("Оба пароля надёжные")
	}
}

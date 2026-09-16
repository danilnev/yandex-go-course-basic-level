package main

import (
	"fmt"
	"unicode"
)

func hasMinimumLength(password string, minLen int) bool {
	return len(password) >= minLen
}

func hasUpper(password string) bool {
	for _, let := range password {
		if unicode.IsUpper(let) {
			return true
		}
	}
	return false
}

func hasLowerCase(password string) bool {
	for _, let := range password {
		if unicode.IsLower(let) {
			return true
		}
	}
	return false
}

func ratePassword(password string) string {
	count := 0
	if hasMinimumLength(password, 8) {
		count++
	}
	if hasUpper(password) {
		count++
	}
	if hasLowerCase(password) {
		count++
	}
	switch count {
	case 3:
		return "Сложный пароль"
	case 2:
		return "Средний пароль"
	case 1:
		return "Слабый пароль"
	default:
		return "Пароль недостаточно безопасен, придумайте новый"
	}
}

func main() {
	fmt.Println(ratePassword("2345"))
	fmt.Println(ratePassword("23453845937954"))
	fmt.Println(ratePassword("2345a"))
	fmt.Println(ratePassword("2345AFf"))
	fmt.Println(ratePassword("2345asdfjs"))
	fmt.Println(ratePassword("2345FF"))
	fmt.Println(ratePassword("2345JFDJSLJF"))
	fmt.Println(ratePassword("2345sdfFKJKD"))
	fmt.Println(ratePassword("aF"))
}

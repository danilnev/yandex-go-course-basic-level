package main

import (
	"fmt"
	"unicode"
)

func checkPassword(password string) bool {
	return hasMinimumLength(password, 8) && hasUpper(password)
}

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

func main() {
	fmt.Println(checkPassword("123a"))
	fmt.Println(checkPassword("123aB"))
	fmt.Println(checkPassword("asdfasdfasfFDFDKSK123a"))
	fmt.Println(checkPassword("123asdfasgfdfgjldsjgkladslkfa"))
}

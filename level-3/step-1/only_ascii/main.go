package main

import (
	"unicode"
)

func CheckOnlyASCII(s string) bool {
	for _, symb := range s {
		if symb > unicode.MaxASCII {
			return false
		}
	}
	return true
}

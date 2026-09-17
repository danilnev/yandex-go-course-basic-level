package main

import (
	"fmt"
	"strings"
)

func NumbersToLetters(input string) string {
	result := strings.ReplaceAll(input, "0", "ноль")
	result = strings.ReplaceAll(result, "1", "один")
	result = strings.ReplaceAll(result, "2", "два")
	result = strings.ReplaceAll(result, "3", "три")
	result = strings.ReplaceAll(result, "4", "четыре")
	result = strings.ReplaceAll(result, "5", "пять")
	result = strings.ReplaceAll(result, "6", "шесть")
	result = strings.ReplaceAll(result, "7", "семь")
	result = strings.ReplaceAll(result, "8", "восемь")
	result = strings.ReplaceAll(result, "9", "девять")
	result = strings.ReplaceAll(result, "+", "плюс")
	result = strings.ReplaceAll(result, "-", "минус")
	result = strings.ReplaceAll(result, "/", "разделить на")
	result = strings.ReplaceAll(result, "*", "умножить на")

	return result
}

func main() {
	fmt.Println(NumbersToLetters("(1 + 2) * 3 / 8"))
}

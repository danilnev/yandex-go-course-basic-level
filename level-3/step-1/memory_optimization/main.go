package main

import (
	"fmt"
	"unicode/utf8"
)

func CountLengthAndBytes(first, second string) string {
	joined := first + second
	num_bytes := len(joined)
	num_runes := utf8.RuneCountInString(joined)
	return fmt.Sprintf(
		"Объединённая строка: %s. "+
			"Количество байт: %d. Количество символов: %d.",
		joined, num_bytes, num_runes,
	)
}

package main

import (
	"fmt"
	"strings"
)

func CheckLetters(text string) string {
	count := strings.Count(text, "е")
	if count == 0 {
		return "Текст готов к публикации!"
	}
	return fmt.Sprintf("Количество возможных ошибок: %d, перепроверьте текст", count)
}

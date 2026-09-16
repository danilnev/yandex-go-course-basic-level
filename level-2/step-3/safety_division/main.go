package main

import (
	"errors"
	"fmt"
)

var (
	DivisionByZeroError = errors.New("division by zero is not allowed")
)

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0.0, DivisionByZeroError
	}
	return float64(a) / float64(b), nil
}

func main() {
	var a, b int
	fmt.Scanln(&a, &b)

	fmt.Println(Divide(a, b))
}

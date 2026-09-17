package main

import "fmt"

func PrettyArrayOutput(array [9]string) {
	for i := 0; i < 9; i++ {
		if i < 7 {
			fmt.Printf("%d я уже сделал: %s\n", i+1, array[i])
		} else {
			fmt.Printf("%d не успел сделать: %s\n", i+1, array[i])
		}
	}
}

package main

import (
	"fmt"
)

func main() {
	var command string
	var num int
	place1, place2, place3, place4, place5 := "-", "-", "-", "-", "-"

	fmt.Scan(&command)
	for command != "конец" {
		switch command {

		case "":
			fmt.Println("")
			fmt.Scan(&command)
			continue

		case "очередь":
			fmt.Printf("1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n", place1, place2, place3, place4, place5)
			fmt.Scan(&command)
			continue

		case "количество":
			free := 0
			if place1 == "-" {
				free++
			}
			if place2 == "-" {
				free++
			}
			if place3 == "-" {
				free++
			}
			if place4 == "-" {
				free++
			}
			if place5 == "-" {
				free++
			}
			fmt.Printf("Осталось свободных мест: %d\nВсего человек в очереди: %d\n", free, 5-free)
			fmt.Scan(&command)
			continue

		case "конец":
			break

		}

		fmt.Scanln(&num)
		if num < 1 || num > 5 {
			fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", num)
			fmt.Scan(&command)
			continue
		}
		if place1 != "-" && place2 != "-" && place3 != "-" && place4 != "-" && place5 != "-" {
			fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", num)
			fmt.Scan(&command)
			continue
		}
		switch num {
		case 1:
			if place1 != "-" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
				fmt.Scan(&command)
				continue
			}
			place1 = command
		case 2:
			if place2 != "-" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
				fmt.Scan(&command)
				continue
			}
			place2 = command
		case 3:
			if place3 != "-" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
				fmt.Scan(&command)
				continue
			}
			place3 = command
		case 4:
			if place4 != "-" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
				fmt.Scan(&command)
				continue
			}
			place4 = command
		case 5:
			if place5 != "-" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
				fmt.Scan(&command)
				continue
			}
			place5 = command
		}

		fmt.Scan(&command)
	}

	fmt.Printf("1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n", place1, place2, place3, place4, place5)
}

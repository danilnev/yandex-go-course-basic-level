package main

import "fmt"

func SwapKeysAndValues(m map[string]string) map[string]string {
	swapped_map := map[string]string{}

	for key, value := range m {
		swapped_map[value] = key
	}

	return swapped_map
}

func main() {
	fmt.Println(SwapKeysAndValues(map[string]string{"sldjflsjdf": "ok", "sld": "no"}))
}

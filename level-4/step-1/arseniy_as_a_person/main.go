package main

import (
	"fmt"
)

type Person struct {
	name    string
	age     int
	address string
}

func (person Person) PrettyPrint() {
	fmt.Printf(
		"Name: %s\n"+
			"Age: %d\n"+
			"Address: %s\n",
		person.name,
		person.age,
		person.address,
	)
}

func main() {
	person := Person{name: "Danil", age: 18, address: "Ryazan"}
	person.PrettyPrint()
}

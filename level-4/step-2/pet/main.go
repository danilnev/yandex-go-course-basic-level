package main

import (
	"fmt"
)

type Animal interface {
	MakeSound()
}

type Dog struct{}

func (d Dog) MakeSound() {
	fmt.Println("Гав!")
}

type Cat struct{}

func (c Cat) MakeSound() {
	fmt.Println("Мяу!")
}

func main() {
	var dog Animal = Dog{}
	var cat Animal = Cat{}
	dog.MakeSound()
	cat.MakeSound()
}

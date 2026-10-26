package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name     string
	Age      int
	IsActive bool
}

func NewUser(name string, age int) (*User, error) {
	if name == "" {
		return nil, errors.New("name is empty for user")
	}
	if age == 0 {
		age = 18
	}
	return &User{Name: name, Age: age, IsActive: true}, nil
}

func main() {
	user1, ok1 := NewUser("", 15)
	user2, ok2 := NewUser("Danil", 0)
	user3, ok3 := NewUser("Pasha", 20)
	fmt.Println(user1, ok1)
	fmt.Println(user2, ok2)
	fmt.Println(user3, ok3)
}

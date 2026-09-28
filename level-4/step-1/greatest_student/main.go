package main

import (
	"fmt"
)

type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (student Student) IsExcellentStudent() bool {
	return student.scoreForOneTask*float64(student.solvedProblems) >= student.passingScore
}

func main() {
	student1 := Student{name: "Danil", solvedProblems: 6, scoreForOneTask: 4.5, passingScore: 20.5}
	student2 := Student{name: "Pasha", solvedProblems: 5, scoreForOneTask: 3.9, passingScore: 20.5}
	fmt.Println(student1.IsExcellentStudent())
	fmt.Println(student2.IsExcellentStudent())
}

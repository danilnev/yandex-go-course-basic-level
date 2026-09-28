package main

import "fmt"

type Employee struct {
	employee string
	position string
	salary   float64
	bonus    float64
}

func (employeer Employee) CalculateTotalSalary() {
	totalSalary := employeer.salary + employeer.bonus

	fmt.Printf(
		"Employee: %s\n"+
			"Position: %s\n"+
			"Total Salary: %.2f\n",
		employeer.employee,
		employeer.position,
		totalSalary,
	)
}

func main() {
	employeer := Employee{employee: "Danil", position: "backend developer", salary: 55.9999, bonus: 10.5555}
	employeer.CalculateTotalSalary()
}

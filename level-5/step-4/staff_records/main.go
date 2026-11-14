package main

import (
	"fmt"
	"slices"
)

type Company struct {
	Workers []Worker
}

type Worker struct {
	Name       string
	Position   string
	Salary     uint
	Experience uint
}

type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	c.Workers = append(c.Workers, Worker{Name: name, Position: position, Salary: salary, Experience: experience})

	return nil
}

var PositionsOrder = map[string]int{
	"директор":       1,
	"зам. директора": 2,
	"начальник цеха": 3,
	"мастер":         4,
	"рабочий":        5,
}

func (c *Company) SortWorkers() ([]string, error) {
	slices.SortFunc(c.Workers, func(a, b Worker) int {
		if a.Salary*a.Experience == b.Salary*b.Experience {
			if PositionsOrder[a.Position] < PositionsOrder[b.Position] {
				return -1
			} else if PositionsOrder[a.Position] > PositionsOrder[b.Position] {
				return 1
			} else {
				return 0
			}
		} else if a.Salary*a.Experience > b.Salary*b.Experience {
			return -1
		} else {
			return 1
		}
	})
	resultStrings := []string{}
	for _, worker := range c.Workers {
		resultStrings = append(resultStrings, fmt.Sprintf("%s - %d - %s", worker.Name, worker.Salary*worker.Experience, worker.Position))
	}
	return resultStrings, nil
}

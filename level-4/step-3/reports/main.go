package main

import (
	"fmt"
	"time"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	ReportID int
	Date     string
	User
}

func CreateReport(user User, reportDate string) Report {
	timeNow := time.Now()
	reportID := timeNow.Second() + timeNow.Nanosecond()
	return Report{ReportID: reportID, Date: reportDate, User: user}
}

func GenerateUserReports(users []User, reportDate string) []Report {
	resultReports := []Report{}
	for _, user := range users {
		newReport := CreateReport(user, reportDate)
		resultReports = append(resultReports, newReport)
	}
	return resultReports
}

func main() {
	userList := []User{
		{ID: 1, Name: "Danil", Email: "email1", Age: 18},
		{ID: 2, Name: "Pasha", Email: "email2", Age: 22},
		{ID: 3, Name: "Alex", Email: "email3", Age: 25},
	}

	userReports := GenerateUserReports(userList, time.Now().Format("02.01.2006 15:04:05"))

	fmt.Println(userReports)
}

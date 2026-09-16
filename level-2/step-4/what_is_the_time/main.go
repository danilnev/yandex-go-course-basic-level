package main

import (
	"errors"
	"strings"
	"time"
)

var (
	DayOrNightError = errors.New("исправь свой ответ, а лучше ложись поспать")
)

func currentDayOfTheWeek() string {
	day := time.Now().Weekday()
	switch day {
	case time.Monday:
		return "Понедельник"
	case time.Tuesday:
		return "Вторник"
	case time.Wednesday:
		return "Среда"
	case time.Thursday:
		return "Четверг"
	case time.Friday:
		return "Пятница"
	case time.Saturday:
		return "Суббота"
	case time.Sunday:
		return "Воскресенье"
	}
	return "Неизвестный день"
}

func dayOrNight() string {
	hour := time.Now().Hour()
	if hour >= 10 && hour < 22 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	weekday := time.Now().Weekday()
	days := int(time.Friday - weekday)
	if days <= 0 {
		days += 7
	}
	return days
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	weekday := currentDayOfTheWeek()
	return strings.EqualFold(weekday, answer)
}

func CheckNowDayOrNight(answer string) (bool, error) {
	current := dayOrNight()
	if len(answer) != 8 {
		return false, DayOrNightError
	}
	return strings.EqualFold(current, answer), nil
}

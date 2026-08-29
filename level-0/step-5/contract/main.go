package main

import (
	"fmt"
	"time"
)

func main() {
	var dateStr, firstname, lastname, middlename string
	var firstSum, secondSum, thirdSum float64

	fmt.Scanln(&dateStr)
	fmt.Scanln(&firstname)
	fmt.Scanln(&lastname)
	fmt.Scanln(&middlename)
	fmt.Scanln(&firstSum)
	fmt.Scanln(&secondSum)
	fmt.Scanln(&thirdSum)

	date, err := time.Parse("02.01.2006", dateStr)

	if err != nil {
		fmt.Printf("Error parsing date: %s\n", err.Error())
	}

	date = date.AddDate(0, 0, 15)

	rublesSum := int(firstSum + secondSum + thirdSum)

	kopecksSum := int((firstSum + secondSum + thirdSum - float64(rublesSum)) * 100)

	fmt.Printf(
		"Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\n"+
			"Дата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\n"+
			"Общая сумма выплат составит %d руб. %d коп.\n\n"+
			"С уважением,\nГл. бух. Иванов А.Е.\n",
		lastname, firstname, middlename, date.Format("02.01.2006"), rublesSum, kopecksSum,
	)
}

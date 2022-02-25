package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(dayOfTheWeek(15, 4, 2008))
}

func dayOfTheWeek(day int, month int, year int) string {
	if month == 1 || month == 2 {
		year -= 1
		month += 10
	} else {
		month -= 2
	}

	var (
		k    = float64(day)
		m    = float64(month)
		C    = float64(year / 100)
		Y    = float64(year % 100)
		days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	)

	w := int(k+math.Floor(2.6*m-0.2)-2*C+Y+math.Floor(Y/4)+math.Floor(C/4)) % 7
	if w < 0 {
		w += 7
	}

	return days[w]
}

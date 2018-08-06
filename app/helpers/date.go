package helpers

import "strconv"

func ConvertStringToInt(year, month, day string) (int, int, int) {
	intYear, _ := strconv.Atoi(year)
	intMonth, _ := strconv.Atoi(month)
	intDay, _ := strconv.Atoi(day)

	return intYear, intMonth, intDay
}

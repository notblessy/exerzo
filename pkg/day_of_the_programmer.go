package pkg

import "fmt"

/*
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

func dayOfProgrammer(year int32) string {
	// Write your code here

	var day, month int32

	if year == 1918 {
		day = 26
		month = 9
	} else if isLeapYear(year) {
		day = 12
		month = 9
	} else {
		day = 13
		month = 9
	}

	return fmt.Sprintf("%02d.%02d.%d", day, month, year)

}

func isLeapYear(year int32) bool {
	if year >= 1700 && year <= 1917 {
		return year%4 == 0
	}

	return (year%400 == 0) || (year%4 == 0 && year%100 != 0)
}

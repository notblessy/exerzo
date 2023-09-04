package pkg

import "fmt"

func TimeConversion(s string) string {
	var militaryTime string

	indicator := s[len(s)-2:]

	hourStr := s[:2]
	hour := s[:2]
	if indicator == "AM" {
		if hourStr == "12" {
			militaryTime = "00" + s[2:len(s)-2]
		} else {
			militaryTime = s[:len(s)-2]
		}
	} else if indicator == "PM" {
		if hourStr == "12" {
			militaryTime = s[:len(s)-2]
		} else {
			hourInt := int(hour[0]-'0')*10 + int(hour[1]-'0') + 12
			militaryTime = fmt.Sprintf("%02d%s", hourInt, s[2:len(s)-2])
		}
	} else {
		return "error time format"
	}

	return militaryTime
}

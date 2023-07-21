package settings

import (
	"strings"
	"time"
)

var DATE_LAYOUT string = "2006-01-02"

func DateNow() string {
	return strings.Split(time.Now().String(), " ")[0]
}

func GenerateDateRange(dateIni string, dateEnd string) []string {
	dates := make([]string, 0)

	startDate, _ := time.Parse(DATE_LAYOUT, dateIni)
	endDate, _ := time.Parse(DATE_LAYOUT, dateEnd)

	currentDate := startDate

	for currentDate.Before(endDate) || currentDate.Equal(endDate) {
		dates = append(dates, currentDate.Format(DATE_LAYOUT))
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return dates
}

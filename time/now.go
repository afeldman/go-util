package time_util

import "time"

func GetToday(format string) (todayString string) {
	today := time.Now()
	todayString = today.Format(format)
	return
}

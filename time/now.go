package time

import "time"

func GetToday(format string) (todayString string) {
	today := time.Now()
	todayString = today.Format(format)
	return
}

func ParseTime(timeString, format string) (time.Time, error) {
	return time.Parse(format, timeString)
}

func StringTime(time_ time.Time, format string) string {
	if format == "" {
		format = "2024-07-02T11:48:05Z"
	}
	if time_.IsZero() {
		return time.Now().UTC().Format(format)
	}
	return time_.UTC().Format(format)
}

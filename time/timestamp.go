package time

import "time"

func ParseTimeStamp(timeStamp int64) time.Time {
	return time.Unix(timeStamp, 0)
}

func Timestamp(time_ time.Time) int64 {
	if time_.IsZero() {
		return time.Now().Unix()
	}
	return time_.Unix()
}

func TimestampNano(time_ time.Time) int64 {
	if time_.IsZero() {
		return time.Now().UnixNano()
	}
	return time_.UnixNano()
}

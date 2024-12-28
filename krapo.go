package krapo

import (
	"time"
)

func Time() string {
	currentTime := time.Now()
	return currentTime.Format("20060102")
}

func LTime(day_ago int) string {
	currentTime := time.Now()
	currentTime = currentTime.AddDate(0, 0, -day_ago)
	return currentTime.Format("20060102")
}

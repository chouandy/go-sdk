package time

import "time"

// NextWeekday next weekday
func NextWeekday(date time.Time, weekday time.Weekday) time.Time {
	return date.AddDate(0, 0, int((7+(weekday-date.Weekday()))%7))
}

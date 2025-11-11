package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"

	dt, err := time.Parse(layout, date)
	if err != nil {
		panic("You're out of time!")
	}

	return dt
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	now := time.Now()

	dt, err := time.Parse(layout, date)
	if err != nil {
		panic("You're out of time!")
	}

	return now.After(dt)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"

	dt, err := time.Parse(layout, date)
	if err != nil {
		panic("You're out of time!")
	}

	hour := dt.Hour()
	if hour >= 12 && hour <= 19 {
		return true
	}

	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"

	dt, err := time.Parse(layout, date)
	if err != nil {
		panic("You're out of time!")
	}

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", dt.Weekday(), dt.Month(), dt.Day(), dt.Year(), dt.Hour(), dt.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}

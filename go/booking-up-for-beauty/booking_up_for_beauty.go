package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	parsedTime, _ := time.Parse("1/2/2006 15:04:05", date)
	return parsedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	appointmentTime, _ := time.Parse("January 2, 2006 15:04:05", date)
	currentTime := time.Now()
	return appointmentTime.Before(currentTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsedTime, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return parsedTime.Hour() >= 12 && parsedTime.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appointmentTime := Schedule(date)
	return fmt.Sprintf(
		"You have an appointment on %s, %s %d, %d, at %d:%d.",
		appointmentTime.Weekday(),
		appointmentTime.Month(),
		appointmentTime.Day(),
		appointmentTime.Year(),
		appointmentTime.Hour(),
		appointmentTime.Minute(),
	)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	return time.Date(currentYear, 9, 15, 0, 0, 0, 0, time.UTC)
}

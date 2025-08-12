package booking

import (
    "fmt"
    "time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    t, err := time.Parse("1/2/2006 15:04:05", date)
    if err != nil {
		return time.Time{} // zero time if input invalid
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    t, err := time.Parse("January 2, 2006 15:04:05", date)
    if err != nil {
        return false
    }
	return t.Before(time.Now().UTC())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
    if err != nil {
        return false
    }
    hour := t.Hour()
	return hour >=12 && hour <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    t, err := time.Parse("1/2/2006 15:04:05", date)
     if err != nil {
        return ""
    }
    formattedTime := t.Format("Monday, January 2, 2006, at 15:04")
  	return fmt.Sprintf("You have an appointment on %s.", formattedTime)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().UTC().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}

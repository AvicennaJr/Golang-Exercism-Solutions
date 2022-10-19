 
package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	dateResult, _ := time.Parse("1/2/2006 15:04:05", date)
    
    return dateResult
    
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	dateResult, _ := time.Parse("January 2, 2006 15:04:05", date)
    if dateResult.Before(time.Now()) {
        return true
    } else {
        return false
    }
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    dateResult, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
    if dateResult.Hour() >= 12 && dateResult.Hour() <= 18 {
        return true
    } else {
    	return false
    }
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	dateResult, _ := time.Parse("1/2/2006 15:04:05", date)
    
    return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.", dateResult.Weekday(), dateResult.Month(), dateResult.Day(), dateResult.Year(), dateResult.Hour(), dateResult.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    anniversary, _ := time.Parse("2006-01-02", "2022-09-15")
    return anniversary
}

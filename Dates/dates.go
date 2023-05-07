package main

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	// Get the current date and time
	now := time.Now()

	// Calculate the start and end of the current week
	startOfWeek := now.AddDate(0, 0, -int(now.Weekday())+1)

	// Print the dates of the current week in word format
	var s = ""
	for i := 0; i < 5; i++ {
		date := startOfWeek.AddDate(0, 0, i)
		day := date.Day()
		weekday := date.Weekday()
		month := date.Month()
		year := date.Year()
		suffix := getSuffix(day)
		s = s + fmt.Sprintf("%s, %d%s %s %d\n", weekday, day, suffix, month, year)
	}
	clipboard.WriteAll(s)
	fmt.Println(s)
	fmt.Println("has been copied to your keyboard")
}

// Helper function to get the suffix for a day (e.g. "st", "nd", "rd", or "th")
func getSuffix(day int) string {
	switch day {
	case 1, 21, 31:
		return "st"
	case 2, 22:
		return "nd"
	case 3, 23:
		return "rd"
	default:
		return "th"
	}
}

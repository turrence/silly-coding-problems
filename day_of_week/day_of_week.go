package leetcode

/*
Given a date, return the corresponding day of the week for that date.

The input is given as three integers representing the day, month and year respectively.

Return the answer as one of the following values {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}.


The given dates are valid dates between the years 1971 and 2100.

*/
import "time"

func dayOfTheWeek(day int, month int, year int) string {
	// this is how i would do it in any other situation LOL
	today := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	DayOfWeek := today.Weekday()
	return DayOfWeek.String()

	/*
			currently not working, i'm not sure why
		    // start this list with today's date
		    dayNames := [7]string{"Thursday", "Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday"}
		    daysInMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

		    leapYearOffset := func(year int) int {
		        if (year % 4 == 0 && year % 100 != 0) || year % 400 != 0 {
		            return 1
		        }
		        return 0
		    }

		    daysSinceStart := func(d, m, y int) int {
		        // start is 31 Dec 1970
		        days := 0

		        // years
		        for i := 1971; i < y; i++ {
		            days += 365 + leapYearOffset(y)
		        }

		        // months
		        for i := 1; i < m - 1; i++ {
		            days += daysInMonth[i-1]
		        }

		        // days
		        if month > 1 {
		            days += leapYearOffset(y)
		        }
		        days += d
		        return days
		    }

		    // make sure dayNames starts with the day it is below
		    knownDateDifference := daysSinceStart(27, 8, 2023)
		    d := daysSinceStart(day, month, year)
		    return dayNames[((d - knownDateDifference) % 7 + 7) % 7]
	*/
}

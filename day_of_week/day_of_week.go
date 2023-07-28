package leetcode

/*
Given a date, return the corresponding day of the week for that date.

The input is given as three integers representing the day, month and year respectively.

Return the answer as one of the following values {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}.


The given dates are valid dates between the years 1971 and 2100.

definitely would not be able to do this if asked in an interview bc i can't even understand it when i have the time to look at it

*/
import "time"

func dayOfTheWeek(day int, month int, year int) string {
	// this is how i would do it in any other situation LOL
	today := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	DayOfWeek := today.Weekday()
	return DayOfWeek.String()

	// this solution is kinda hard
	/* revolves around using the knowledge of todays date
	daysInMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	leapYearOffset := func(year int) int {
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
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
		for i := 1; i < m; i++ {
			days += daysInMonth[i-1]
		}

		// days
		if month > 2 {
			days += leapYearOffset(y)
		}
		days += d
		return days
	}

	// make sure this date day is 0-index in dayNames
	knownDate := daysSinceStart(27, 8, 2023)
	d := daysSinceStart(day, month, year)
	// start this list with today's date
	dayNames := [7]string{"Thursday", "Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday"}
	return dayNames[((d-knownDate)%7+7)%7]
	*/
}

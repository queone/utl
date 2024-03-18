package utl

import (
	"fmt"
	"strconv"
	"time"
)

// Return absolute value of int value
func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Returns absolute value of int64 value
func Int64Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// Converts string number to int64
func StringToInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// Converts int64 number to string
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Check if string is a valid date in expectedFormat. Return true if so, false otherwise.
// See https://pkg.go.dev/time
func ValidDate(dateString, expectedFormat string) bool {
	_, err := time.Parse(expectedFormat, dateString)
	return err == nil
}

// Converts an epoch timestamp in int64 format to a time.Time object.
func EpocInt64ToTime(epocInt int64) time.Time {
	return time.Unix(epocInt, 0)
}

// Converts an epoch timestamp in string format to a time.Time object.
// Returns a time.Time object and an error if the conversion fails.
func EpocStringToTime(epocString string) (time.Time, error) {
	epocInt64, err := StringToInt64(epocString)
	return time.Unix(epocInt64, 0), err
}

// Converts dateString from source format to destination format.
// Returns date string in destination format and an error if the conversion fails.
func ConvertDateFormat(dateString, srcFormat, dstFormat string) (string, error) {
	t, err := time.Parse(srcFormat, dateString)
	if err != nil {
		return "", err
	}
	return t.Format(dstFormat), nil
}

// Convert dateString, given in dateFormat, to Unix Epoc seconds int64
func DateStringToEpocInt64(dateString, dateFormat string) (int64, error) {
	t, err := time.Parse(dateFormat, dateString) // First, convert string to Time
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil // Finally, convert Time type to Unix epoc seconds
}

// Print yyyy-mm-dd date for given number of +/- days in future or past
func GetDateInDays(days string) time.Time {
	now := time.Now().Unix()
	daysInt64, err := StringToInt64(days)
	if err != nil {
		panic(err.Error())
	}
	now += (daysInt64 * 86400) // 86400 seconds in a day
	return EpocInt64ToTime(now)
}

// Returns true if given year is a leap year. False otherwise.
func IsLeapYear(year int64) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// Calculate and return number of +/- days from NOW to date given
// Note: Calculations are all in UTC time. And it takes leap year into account.
func GetDaysSinceOrTo(date1 string) int64 {
	start, err := time.Parse("2006-01-02", date1)
	if err != nil {
		panic(err.Error())
	}

	end := time.Now().UTC()

	var days int64 = 0
	var sign int64 = -1

	if start.After(end) {
		start, end = end, start
		sign = 1
	}

	for start.Year() < end.Year() || (start.Year() == end.Year() && start.YearDay() < end.YearDay()) {
		days++
		start = start.AddDate(0, 0, 1)

		// Adjust for leap years
		if start.Month() == time.February && start.Day() == 28 && IsLeapYear(int64(start.Year())) {
			days++
			start = start.AddDate(0, 0, 1)
		}
	}

	return sign * days
}

// Print number of days, also in years and days
func PrintDays(days int64) {
	days_abs := Int64Abs(days)
	var years int64 = 0

	for days_abs >= 365 {
		leap := int64(0)
		if IsLeapYear(years) {
			leap = 1
		}
		if days_abs >= (365 + leap) {
			days_abs -= (365 + leap)
			years++
		} else {
			break
		}
	}

	if years > 0 {
		fmt.Printf("%d (%d years + %d days)\n", days, years, days_abs)
	} else {
		fmt.Println(days)
	}
}

// Return number of days between 2 dates
func GetDaysBetween(date1, date2 string) int64 {
	epoc1, err := DateStringToEpocInt64(date1, "2006-01-02")
	if err != nil {
		panic(err.Error())
	}
	epoc2, err := DateStringToEpocInt64(date2, "2006-01-02")
	if err != nil {
		panic(err.Error())
	}

	return (Int64Abs(epoc1-epoc2) / 86400)
}

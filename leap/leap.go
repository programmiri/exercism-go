// Package leap implements method to detact a leap year.
package leap

// IsLeapYear report if a given year it is a leap year.
func IsLeapYear(year int) bool {
  return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

package exercism

import (
	"errors"
	"fmt"
)

// solves puzzle https://exercism.org/tracks/go/exercises/expenses

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(record Record) bool) []Record {
	var solution []Record
	for _, record := range in {
		if predicate(record) {
			solution = append(solution, record)
		}
	}
	return solution
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		if r.Day >= p.From && r.Day <= p.To {
			return true
		}
		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		if r.Category == c {
			return true
		}
		return false
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var totalByPeriod float64
	records := Filter(in, ByDaysPeriod(p))
	for _, record := range records {
		totalByPeriod += record.Amount
	}
	return totalByPeriod
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, c string, p DaysPeriod) (float64, error) {
	var totalByCategory float64
	records := Filter(in, ByCategory(c))
	if len(records) == 0 {
		err := errors.New(fmt.Sprintf("There is no records for the category %s", c))
		return totalByCategory, err
	}
	records = Filter(records, ByDaysPeriod(p))
	for _, record := range records {
		totalByCategory += record.Amount
	}
	return totalByCategory, nil
}

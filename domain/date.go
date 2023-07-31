package domain

import "time"

type Date struct {
	Month string
	Day   int
}

func NewDateFromTime(time time.Time) (Date, error) {
	return Date{time.Month().String(), time.Day()}, nil
}

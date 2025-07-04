package model

import "time"

type Entry struct {
	Amount   float64
	Type     string
	Category string
	Date     time.Time
}

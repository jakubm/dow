package main

import "time"

func dow(date string) string {
	t, _ := time.Parse("2006-01-02", date)
	return t.Weekday().String()
}

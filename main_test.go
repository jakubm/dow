package main

import "testing"

func TestAbs(t *testing.T) {
	const date = "1978-05-23"
	const expected = "Tuesday"

	got := dow(date)

	if got != expected {
		t.Errorf("Case %s, expected %s, got %s!", date, expected, got)
	}
}

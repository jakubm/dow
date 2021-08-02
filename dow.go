package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	params := os.Args[1:]
	if len(params) > 0 {
		for _, v := range params {
			t, _ := time.Parse("2006-01-02", v)
			fmt.Println(v, t.Weekday())
		}
	} else {
		fmt.Println("dow - program for showing day of week for a date")
		fmt.Println("	Usage: dow YYYY-MM-DD [YYYY-MM-DD, ...]")
	}
}

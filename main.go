package main

import (
	"fmt"
	"os"
)

func main() {
	params := os.Args[1:]
	if len(params) > 0 {
		for _, v := range params {
			fmt.Println(v, dow(v))
		}
	} else {
		fmt.Println("dow - program for showing day of week for a date")
		fmt.Println("	Usage: dow YYYY-MM-DD [YYYY-MM-DD, ...]")
	}
}

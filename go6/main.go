package main

import (
	"fmt"
	"time"
)

func main() {
	level := 1
	switch level {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("GOOD")
	default:
		fmt.Println("QWQ")
	}
	// ***
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("早上")
	default:
		fmt.Println("下午")
	}
}

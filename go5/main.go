package main

import "fmt"

func main() {
	score := 30

	if score >= 30 {
		fmt.Println("MVP")
	} else if score >= 20 {
		fmt.Println("GOOD")
	} else if score >= 10 {
		fmt.Println("OK")
	} else {
		fmt.Println("BAD")
	}
}

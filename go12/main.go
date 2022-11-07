package main

import "fmt"

func sum(num int) int {
	if num == 1 {
		return num
	}

	return sum(num-1) + num
}

func main() {
	fmt.Println(sum(10))
}

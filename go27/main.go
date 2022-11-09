package main

import "fmt"

func intergers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println("主線程開始......")
	myInt := intergers()
	// fmt.Println(result)

	for i := 1; i <= 5; i++ {
		fmt.Println(myInt())
	}
}

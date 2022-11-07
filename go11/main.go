package main

import "fmt"

func plus2(a int, b int) int {
	return a + b
}

func plus3(a, b, c int) int {
	return a + b + c
}

func calABCD(a, b int) (int, int, int, int) {
	return a + b, a - b, a / b, a * b
}

func sum(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func main() {
	fmt.Println(plus2(1, 1))
	fmt.Println(plus3(1, 1, 1))
	a1, a2, a3, a4 := calABCD(1, 2)
	a := [4]int{a1, a2, a3, a4}
	fmt.Println(a[2])
	fmt.Println(a)
	fmt.Println(sum(1, 1, 1, 1))
	fmt.Println(sum(1, 1, 1, 1, 1, 1, 1, 1, 1))
}

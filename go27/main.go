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
	myInt1 := intergers()
	result1 := myInt1()
	result2 := myInt1()
	fmt.Println(result1)
	fmt.Println(result2)

	myInt2 := intergers()
	myInt3 := intergers()
	fmt.Println(myInt2())
	fmt.Println(myInt3())

	// fmt.Println(result)

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(myInt())
	// }
}

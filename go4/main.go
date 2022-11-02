package main

import "fmt"

func main() {
	fmt.Println("------------------------------")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println("------------------------------")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("------------------------------")
	for {
		fmt.Println("123")
		break
	}
	fmt.Println("------------------------------")
	for n := 0; n <= 10; n++ {
		if n%2 == 0 {
			// 偶數繼續
			continue
		}
		fmt.Println(n)
	}
	fmt.Println("------------------------------")
}

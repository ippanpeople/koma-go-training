package main

import "fmt"

func MyDeferFunc1() {
	defer fmt.Println("start")
	fmt.Println("end")
}

// 後進先出
func MyDeferFunc2() {
	defer fmt.Println("處理1")
	defer fmt.Println("處理2")
	defer fmt.Println("處理3")
}

func main() {
	fmt.Println("主線程開始運行...")
	MyDeferFunc1()

	// defer func() {
	// 	fmt.Println("處理1")
	// 	fmt.Println("處理2")
	// 	fmt.Println("處理3")
	// }()

	defer MyDeferFunc2()
}

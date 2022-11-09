package main

import "fmt"

func main() {

	fmt.Println("主線程開始......")
	//標籤
main:
	for {
		for {
			fmt.Println("開始執行")
			break main
		}
		fmt.Println("死循環中")
	}
	fmt.Println("程序結束")

}

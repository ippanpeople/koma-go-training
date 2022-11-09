package main

import "fmt"

func main() {
	fmt.Println("主線程開始運行...")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("別怕")
		}
	}()
	panic("不好有病毒！！")

	fmt.Println("運行中...")
}

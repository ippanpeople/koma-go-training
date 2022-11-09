package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("主線程開始運行...")

	file, err := os.Create("mytext.txt")
	// file.Close()
	defer file.Close()

	if err != nil {
		fmt.Println("文件打開錯誤", err)
	}
	file.WriteString("i love golang !! ")

	// file.Close()
}

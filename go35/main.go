package main

import "fmt"

func init() {
	fmt.Println("初始化函數 init1('系統初始化')")
}
func init() {
	fmt.Println("初始化函數 init2('環境初始化')")
}
func init() {
	fmt.Println("初始化函數 init3('資料庫初始化')")
}

func main() {
	fmt.Println("主線程運行開始...")
}

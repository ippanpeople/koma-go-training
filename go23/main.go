package main

import "fmt"

func main() {
	// 匿名函數, 須以變數復職定義
	f1 := func(x, y int) int {
		return x + y
	}
	result1 := f1(10, 20)
	fmt.Println(result1)

	// 匿名函數, 定義完馬上使用
	result2 := func(a, b int) int {
		return a * b
	}(100, 200)
	fmt.Println(result2)

	//
	fmt.Println(
		func(name string) string {
			return "hello  " + name
		}("rinrin"))
}

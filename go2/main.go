package main

import "fmt"

func main() {
	// 字符型
	var a = "你好GO世界"
	b := "hell go world"
	fmt.Println(a)
	fmt.Println(b)
	// 數值型
	var c, d int = 10, 20
	fmt.Println(c, d)
	// 浮點型
	var e float32
	e = 3.1415926
	fmt.Println(e)
	// 定義加直接付值
	f := "123"
	fmt.Println(f)
	// 指針類型
	var pf *string
	pf = &f
	fmt.Println(pf)
	fmt.Println(*pf)
	fmt.Println(&f)
}

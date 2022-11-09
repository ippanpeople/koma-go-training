// 空街口的類型轉換
package main

import "fmt"

func myfunc(val interface{}) {
	fmt.Println(val)

	switch val.(type) {
	case bool:
		fmt.Println("布爾型")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("default")

	}
}

func main() {
	fmt.Println("主線程運行開始......")

	myfunc(123)
	myfunc("abc")
	myfunc(true)
	myfunc(3.14323452)

	var x interface{} = 100
	i1 := x.(int)
	fmt.Println(i1)
	fmt.Println(i1 + 100)

	x = 3.14
	i2 := x.(float64)
	fmt.Println(i2)
	fmt.Println(i2 + 100)

	i3, isFloat64 := x.(float64)
	fmt.Println(isFloat64)
	fmt.Println(i3)
}

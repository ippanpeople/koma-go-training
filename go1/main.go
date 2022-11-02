// 包名
package main

//	如果定義後沒使用則會報錯
import "fmt"

// 主入口函數
func main() {
	// 字符串
	fmt.Println("go" + "lang")
	// 數值
	fmt.Println("1 + 1 = ", 1+1)
	// 小數
	fmt.Println("10.0 / 2,0 = ", 10.0/2.0)
	// 布林值
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(!true)
	// 指針
	a := 5
	// int類型的指針
	var pointer_a *int
	//	獲取a的地址
	pointer_a = &a
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(pointer_a)
	fmt.Println(&pointer_a)
	fmt.Println(*pointer_a)

}

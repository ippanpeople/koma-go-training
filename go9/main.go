// 哈希字典 ＝ key:value 的無序數組
package main

import "fmt"

func main() {
	// 定義一個哈希字典
	m := make(map[string]int) //make分配內存 [string]key類型 int值類型
	// x := map[string]int //定義一個空字典存的話沒有賦值且沒有預先分配內的話會報錯
	// 給哈希字典賦值
	m["key1"] = 1
	m["key2"] = 2

	fmt.Println(m)
	fmt.Println(m["key1"])
	fmt.Println(len(m))

	// 定義家賦值一個哈希字典
	n := map[string]int{"key1": 1, "key2": 2}
	fmt.Println(n)

	// 給變量賦值
	v := m["key2"]
	fmt.Println(v)

	// 刪除一個哈希值
	delete(m, "key1")
	fmt.Println(m)

}

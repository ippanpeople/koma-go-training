// range : 用以迭代數組, 切片, 字典, 字符串, 通道的函數
package main

import "fmt"

func main() {
	// 迭代數組
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	// 連帶索引一起迭代
	for n, num := range nums {
		fmt.Println("index : ", n, ", num : ", num)
	}

	//	迭代字典
	maps := map[string]string{"k1": "apple", "k2": "banana"}
	for k, v := range maps {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// 僅迭代key
	for k := range maps {
		fmt.Println("key : ", k)
	}

	// 迭代字符串
	for i, c := range "i love u" {
		fmt.Println("word", i, " : ", string(c))
	}
}

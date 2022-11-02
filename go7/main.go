package main

import "fmt"

func main() {
	//	數組的定義, 為定義前初始值都是0
	var a [5]int
	fmt.Println("數組 : ", a)
	a[4] = 100
	fmt.Println("數組 : ", a[4])
	fmt.Println("數組 : ", a)
	fmt.Println("數組 : ", a[4:])
	// 數組的長度
	fmt.Println(len(a))
	// 數組加付值
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("數組 :", b)
	// 二維數組
	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println(c)
}

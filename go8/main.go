package main

import "fmt"

func main() {
	//	創建空切片
	s := make([]string, 3)
	fmt.Println(s)
	fmt.Println(len(s))
	// 切片賦值
	s[0] = "1"
	s[1] = "2"
	s[2] = "3"
	fmt.Println(s)
	// 內容追加
	s = append(s, "4")
	fmt.Println(s)
	// 創建新的切片
	c := make([]string, len(s))
	// 拷貝就切片內容到新切片
	copy(c, s)
	fmt.Println(len(c))
	fmt.Println(c)
	var pc *[]string
	pc = &c
	fmt.Println(pc)
	fmt.Println(*pc)
	//	切下切片
	l := s[2:4]
	fmt.Println(l)
	// 創建數組
	t := []string{"1", "2", "3", "4", "5"}
	// 數組切片
	fmt.Println(t[2:4])
	var pt *[]string
	pt = &t
	fmt.Println(pt)
	fmt.Println(*pt)
}

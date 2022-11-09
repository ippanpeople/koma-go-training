package main

import "fmt"

func MyLaterFunc() func(string) string {
	var _oldVal string
	return func(newVal string) string {
		result := _oldVal
		_oldVal = newVal
		return result
	}
}

func main() {
	fmt.Println("主線程開始......")
	f1 := MyLaterFunc()
	result := f1("hello")
	fmt.Println("1. ", result)

	result = f1("go")
	fmt.Println("2. ", result)
}

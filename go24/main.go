package main

import "fmt"

// 定義一個可以返回函數類型的函數
func returnDatabaseType() func() {
	return func() {
		fmt.Println("RDB(PostgreSQL)")
	}
}

func fetchDatabaseType() func(string) string {
	return func(name string) string {
		return "Datebase is " + name
	}
}

func main() {
	resultFunc := returnDatabaseType()
	resultFunc()

	resultFunc2 := fetchDatabaseType()
	result := resultFunc2("MySQL")
	fmt.Println(result)
}

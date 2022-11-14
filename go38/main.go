package main

import (
	"fmt"

	// "db"
	// DB "db"
	. "db"
)

func main() {
	fmt.Println("主線程開始...")

	// 包引用
	// fmt.Println(db.Connection)
	// db.ConnectDB()
	// db.CloseDB()

	// 自定義前綴
	// DB.ConnectDB()
	// DB.CloseDB()

	// 省略前綴的寫法（不推薦）
	ConnectDB()
	CloseDB()
}

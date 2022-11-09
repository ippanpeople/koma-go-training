// callback function
package main

import (
	"fmt"
	"time"
)

func updateDB(cb func()) {
	fmt.Println("數據處理中...")
	time.Sleep(time.Second)
	fmt.Println("數據庫更新完成, 調用回調函數通知主線程.")
	cb()
}

func main() {
	fmt.Println("主線程開始運行...")

	updateDB(func() {
		fmt.Println("->", "收到回調函數訊息, 數據庫處理正常結束")
	})

	fmt.Println("主現成運行結束.")
}

// 通道 - 多協程之間的通信
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 定義一個字符型的通道
	message := make(chan string)
	go func() {
		for i := 1; i <= 3; i++ {
			message <- (strconv.Itoa(i) + ".hello channel.")
		}
	}()

	//	接收通道發送的消息, 收到消息主線程才會繼續執行
	receive := ""
	receive = <-message
	fmt.Println(receive)
	receive = <-message
	fmt.Println(receive)
	receive = <-message
	fmt.Println(receive)
}

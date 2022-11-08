// 帶緩衝區的通道
package main

import "fmt"

func main() {
	// 一般沒有緩衝區的情況下, 發一條消息在未被接收的情況下
	// 線程即阻塞鎖死
	// 建立一個有三個緩衝區的通道,
	// 亦即緩衝區消息達到三個時會組塞當前線程
	message := make(chan string, 3)

	// 發送消息
	message <- "消息 1"
	message <- "消息 2"
	// xxxxxx 當沒有其他協程接收訊息
	// 且緩衝區只有2的時候, 主線程將在發送第三個消息前
	// 主線程被阻塞鎖死
	message <- "消息 3"

	// 收消息
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}

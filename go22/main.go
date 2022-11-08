// 通道同步
// 通道處理 等待其他併發線程的結束
package main

import (
	"fmt"
	"time"
)

func action(done chan bool) {
	fmt.Println("Loading......")
	time.Sleep(2 * time.Second)
	fmt.Println("Finished")
	// 邏輯處理完畢後, 發送通道消息
	done <- true
}

func main() {
	done := make(chan bool)

	go action(done)

	//等待通道消息
	<-done
}

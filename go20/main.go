// 循環接收通道的訊息
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 定義一個字符串型的通道
	message := make(chan string)

	go func() {
		for i := 1; i <= 5; i++ {
			if i == 5 {
				message <- ""
			} else {
				message <- (strconv.Itoa(i) + ".hello channel.")
			}
		}
	}()

	for receive := range message {
		if receive == "" {
			break
		} else {
			fmt.Println(receive)
		}
	}

	// receive := ""
	// receive = <-message

	// for receive != "" {
	// 	fmt.Println(receive)
	// 	receive = <-message
	// }
}

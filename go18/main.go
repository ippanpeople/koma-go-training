// 併發編成 - 協程
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func sayHello(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println("hello ", name, " : ", i)
	}
}
func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	// 同步執行函數
	// sayHello("rinriN")
	// sayHello("rinriNN")
	// sayHello("rinriNNN")

	// 異步執行函數
	go sayHello("rin")
	go sayHello("rinrin")
	go sayHello("lin")

	// 匿名函數異步執行
	go func(msg string) {
		fmt.Println("this is a", msg)
	}("lesson")

	// 等待1秒鐘 等異步函數執行完
	time.Sleep(time.Second)
}

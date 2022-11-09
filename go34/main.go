package main

import (
	"fmt"
	"time"
)

func sayHello(id string) {
	i := 1
	for {
		fmt.Println("hello", id, ":", i)
		i++
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("主線程開始運行...")
	go sayHello("1")
	go sayHello("2")

	i := 100
	for {
		fmt.Println("Hello MainGo", ":", i)
		i++
		time.Sleep(time.Second)
	}
}

package main

import (
	"fmt"
	"time"
)

func worker(p_name string, p_chan chan int) {
	for {
		val, ok := <-p_chan
		if ok {
			fmt.Println(p_name, val, ok)
		} else {
			break
		}
	}
}

func main() {
	fmt.Println("主線程開始運行...")

	chan1 := make(chan int)
	chan2 := make(chan int)
	defer close(chan1)
	defer close(chan2)

	go worker("player1", chan1)
	go worker("player2", chan2)

	i := 1

	for i <= 5 {
		chan1 <- i
		chan2 <- i * 10
		i++
		time.Sleep(time.Second)
	}

	fmt.Println("主線程結束運行...")
	time.Sleep(time.Second)
}

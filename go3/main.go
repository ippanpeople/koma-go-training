package main

import (
	"fmt"
	"math"
)

const message string = "gogogo"

func main() {
	fmt.Println(message)
	// const 付值會報錯
	// message = "nonono"

	// go內置的定數
	fmt.Println(math.Pi)
	fmt.Println(math.Cos(math.Pi / 6))
}

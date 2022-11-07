package main

import "fmt"

func byValue(ival int) {
	ival = 0
}

func byPointer(iptr *int) {
	*iptr = 0
}

func main() {
	i := 10
	fmt.Println("init:", i)
	byValue(i)
	fmt.Println("byval:", i)

	fmt.Println("i address:", &i)
	byPointer(&i)
	fmt.Println("bypointer:", i)
	fmt.Println("i address:", &i)

}

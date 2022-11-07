package main

import "fmt"

type user struct {
	name     string
	password string
	age      int
}

func main() {
	//	實例化 user 結構體 user1
	user1 := user{name: "koma", password: "123", age: 20}
	fmt.Println(user1)
	fmt.Println(user1.name, user1.password, user1.age)
	// 聲明一個user1的指針
	user1p := &user1
	fmt.Println(user1p)
	fmt.Println(*user1p)
	fmt.Println(user1p.name, user1p.password, user1p.age)

	// 利用指針(user1p)給 user1賦值, 會改變user1的內容,
	// 因為user1p 與 user1 共用內存
	user1p.name = "jerry"
	user1p.password = "456"
	user1p.age = 21
	fmt.Println(user1)
	fmt.Println(*user1p)

}

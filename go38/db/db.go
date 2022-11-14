package db

import "fmt"

const (
	Connection string = "12.0.0.1:3306"
)

func ConnectDB() {
	fmt.Println("DB connecting...")
}

func CloseDB() {
	fmt.Println("DB close...")
}

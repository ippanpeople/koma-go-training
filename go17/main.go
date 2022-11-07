package main

import (
	"errors"
	"fmt"
)

// 帶錯誤返回值的函數
func f1(num int) (int, error) {
	if num < 0 {
		return -1, errors.New("參數錯誤！！")
	}
	return 2 * num, nil
}

// 自定義錯誤結構體
type argError struct {
	arg int
	msg string
}

// 定義錯誤結構體函數
func (e *argError) Error() string {
	return fmt.Sprintf("%d -> %s", e.arg, e.msg)
}
func f2(num int) (int, error) {
	if num < 0 {
		return -1, &argError{num, "參數不能為負"}
	}
	return 2 * num, nil
}
func main() {
	result, err := f1(10)
	fmt.Println(result, err)
	result, err = f1(-10)
	fmt.Println(result, err)

	result, err = f2(10)
	fmt.Println(result, err)
	result, err = f2(-10)
	fmt.Println(result, err)

}

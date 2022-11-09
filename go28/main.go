package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("主線程運行開使......")

	var strAge string = "abc"
	var intAge, err = strconv.Atoi(strAge)

	if err != nil {
		fmt.Println("出錯啦, 類型沒轉換成功:", err)
	} else {
		fmt.Println("永遠的", intAge, "歲")
	}

}

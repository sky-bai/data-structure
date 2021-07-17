package main

import "fmt"

var intb chan string

func main() {
	intb = make(chan string, 2)
	intb <- "bai"
	intb <- "wxy"

	fmt.Println(<-intb)
	// 这tm就只能取一个
	fmt.Println(<-intb)
	close(intb)
	intb <- "baiwxy"

}

// 数组不是引用类型 通过make创建

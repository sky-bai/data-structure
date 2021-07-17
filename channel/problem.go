package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义方向 往外读
	var a <-chan time.Time

	t := time.Now()
	fmt.Println(t)
	a = time.After(2 * time.Second)

	select {
	case d := <-a:
		fmt.Println(d)
	}
}

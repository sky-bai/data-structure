package main

var inttype chan int
var m map[string]chan string

func main() {
	// make 返回类型 new 返回指针
	inttype = make(chan int)

}

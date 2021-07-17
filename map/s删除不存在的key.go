package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	a1 := a[3:4:4]
	fmt.Println(a1[0])
}

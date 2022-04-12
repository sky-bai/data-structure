package main

import "fmt"

func main() {
	deferFuncParameter()
}
func deferFuncParameter() {
	var aInt = 1
	defer fmt.Println(aInt)

	aInt = 2
	defer fmt.Println(aInt)

	aInt = 3
	defer fmt.Println(aInt)

	return
}

type ll struct {
}

func (l *ll) get() {

}

func (l ll) gaet() {

}

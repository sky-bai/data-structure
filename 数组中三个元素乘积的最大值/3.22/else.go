package main

import "fmt"

func main() {
	pan(3)
}

func pan(a int) {

	if a > 1 {
		fmt.Println("---", a)
	} else if a > 2 {
		fmt.Println("ppp", a)
	}

}

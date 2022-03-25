package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0}
	a2 := []int{6, 7, 8, 9, 10}

	top := 0
	down := 0

	//i := 0
	for i := 0; i < len(a1); i++ {

		if a1[top] > a2[down] {
			a1[i], a2[down] = a2[down], a1[i]
			top++

		} else if a1[top] != 0 && a1[top] <= a2[down] {
			top++
		} else if a1[top] == 0 {
			a1[i] = a2[down]
			top++
			down++
		}
	}

	fmt.Println(a1)

}

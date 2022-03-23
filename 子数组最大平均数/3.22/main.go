package main

import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 89, 9}
	k := 3

	max := 0

	sum := 0
	for i := 0; i < k; i++ {
		sum += num[i]
	}

	for i := k; i < len(num); i++ {
		sum = sum + num[i] - num[i-k]
		if sum > max {
			max = sum
		}
	}

	fmt.Println(max)
}

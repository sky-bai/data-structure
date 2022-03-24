package main

import "fmt"

func primer(n int) int {

	nums := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if !nums[i] {
			count++
			for j := 2 * i; j < n; j += i {
				nums[j] = true
			}
		}

	}
	return count
}

func main() {
	num := primer(100)
	fmt.Println(num)
}

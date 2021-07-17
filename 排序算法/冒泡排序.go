package main

import "fmt"

func (c *com) max(nums []int) int {
	left := len(nums) - 2
	right := len(nums) - 1
	for nums[left] <= nums[right] {
		nums[left], nums[right] = nums[right], nums[left]
		if left == 0 {
			break
		}
		left--
		right--
	}
	return nums[left]
}

type com struct {
}

func main() {
	m2 := []int{1, 2, 8, 4, 5}
	com1 := com{}
	max := com1.Maopao(m2)
	fmt.Printf("最大：%d\n", max)
}

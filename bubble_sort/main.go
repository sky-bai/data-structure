package main

import "fmt"

// 冒泡排序
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)

func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ { // 确定已经排序过的位置
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums
}

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(BubbleSort(arr))
}

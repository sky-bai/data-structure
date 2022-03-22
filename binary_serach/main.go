package main

func main() {

}

// 非递归实现
func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// 最大堆
// 堆的定义：
// 堆是一个完全二叉树，其中每个结点的值都大于或等于其左右孩子结点的值。
// 堆的特点：
// 1. 堆中的结点是从1开始编号的，即第一个结点的编号为1，最后一个结点的编号为n。
// 2. 堆中每个结点的值都大于或等于其左右孩子结点的值。
// 3. 堆中每个结点的值都大于或等于其父结点的值。
// 4. 堆中每个结点的值都大于或等于其祖父结点的值。
// help me to create a max heap function
func maxHeap(nums []int, i int) {
	left := 2 * i
	right := 2*i + 1
	max := i
	if left <= len(nums) && nums[left] > nums[max] {
		max = left
	}
	if right <= len(nums) && nums[right] > nums[max] {
		max = right
	}
	if max != i {
		nums[i], nums[max] = nums[max], nums[i]
		maxHeap(nums, max)
	}
}

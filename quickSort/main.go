package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	low := make([]int, 0)
	high := make([]int, 0)
	mid := make([]int, 0)

	spcData := arr[0]
	mid = append(mid, spcData)

	for i := 1; i < len(arr); i++ {
		if arr[i] < spcData {
			low = append(low, arr[i])
		} else if arr[i] > spcData {
			high = append(high, arr[i])
		} else {
			mid = append(mid, arr[i])
		}

	}
	// 将数据分开 每一次分割数据 然后组合数据
	low, high = quickSort(low), quickSort(high)
	myArr := append(append(low, mid...), high...)
	return myArr
}

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(quickSort(arr))
}

// 从一组数据中找到一个基准值 将大于这个数的值放入一个数组 小于这个数的值放入另一个数组

// 1, 9, 10, 30, 2, 5, 45
// low = 0 high = 9 10 30 2 5 45 mid = 1
// 9 10 30 2 5 45
// mid = 9  low = 2 5  high = 10 30 45
// low = 2 5 mid = 2 high = 5
//

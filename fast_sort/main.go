package main

// 将数组分为三份 找出一个值  high mid low 三组

func Fast_Sort(a []int) []int {

	if len(a) <= 1 {
		return a
	}

	data := a[0]

	low := make([]int, 0)
	mid := make([]int, 0)
	high := make([]int, 0)

	for i := 1; i < len(a); i++ {
		if a[i] > data {
			high = append(high, a[i])
		} else if a[i] < data {
			low = append(low, a[i])
		} else {
			mid = append(mid, a[i])
		}
	}

	low, high = Fast_Sort(low), Fast_Sort(high)

	myData := append(append(low, mid...), high...)
	return myData

}

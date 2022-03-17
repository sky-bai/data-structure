package main

func sort(a []int) []int {
	// 最后一次对传入参数进行判断 找到最后一次的结果 这里对入参和出参做了结果 最后这一次的入参结果是好多给了明确的结果
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

	low, high = sort(low), sort(high)

	myData := append(append(low, mid...), high...)
	return myData

}

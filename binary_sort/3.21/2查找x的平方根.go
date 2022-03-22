package main

func main() {

}

func bin_sort(a int) int {
	index, low, high := -1, 0, a

	mid := (low + high) / 2
	for low <= high {
		if mid*mid < a {
			index = mid
			low = mid + 1
		} else if mid*mid > a {
			high = mid - 1
		} else {
			return mid
		}
	}

	return index
}

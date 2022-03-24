package main

func mySqrt(a int) int {

	a1 := float64(a)
	return int(sort(a1, a1))

}

func sort(x float64, a float64) float64 {

	res := (a/x + x) / 2

	if res == x {
		return res
	} else {
		return sort(x, a)
	}

}

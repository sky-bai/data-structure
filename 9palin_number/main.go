package _palin_number

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	if x < revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 偶 / 奇
	return x == revertedNumber || x == revertedNumber/10

}

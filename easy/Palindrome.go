package easy

// IsPalindrome 回文数
func IsPalindrome(x int) bool {
	y := 0
	n := x
	z := 0
	for x > 0 {
		y = x % 10
		z = (z + y) * 10
		x = x / 10
	}
	return z/10 == n
}

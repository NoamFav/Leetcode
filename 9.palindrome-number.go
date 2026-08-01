package leetcode

// @leet start
func isPalindrome(x int) bool {
	var rev int
	var digit int
	var number int

	number = x
	for number > 0 {
		digit = number % 10
		rev = rev*10 + digit
		number /= 10
	}
	return x == rev

}

// @leet end

package leetcode

// @leet start
import (
	"math"
)

func myAtoi(s string) int {
	i, sign, result := 0, 1, 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		result = result*10 + int(s[i]-'0')
		if sign*result <= math.MinInt32 {
			return math.MinInt32
		}
		if sign*result >= math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * result
}

// @leet end

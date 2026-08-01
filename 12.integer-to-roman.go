package leetcode

// @leet start
func intToRoman(num int) string {
	val := map[int]string{
		1: "I", 4: "IV", 5: "V", 9: "IX",
		10: "X", 40: "XL", 50: "L", 90: "XC",
		100: "C", 400: "CD", 500: "D", 900: "CM",
		1000: "M",
	}
	var result string

	getpower := func(power int) {
		digit := num / power
		num %= power
		switch digit {
		case 4:
			result += val[digit*power]
		case 9:
			result += val[digit*power]
		default:
			if digit >= 5 {
				result += val[5*power]
				digit -= 5
			}
			for range digit {
				result += val[power]
			}
		}
	}

	for i := 1000; i >= 1; i /= 10 {
		getpower(i)
	}

	return result
}

// @leet end

package leetcode

// @leet start
func romanToInt(s string) int {
	val := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	res := 0
	num := 0
	current := 0
	next := 0

	length := len(s)
	for i := 0; i < length; i++ {
		current = val[s[i]]
		next = 0
		num = 0

		if i < length-1 {
			next = val[s[i+1]]
		}
		isvalid := current*10 == next || current*5 == next
		if current < next {
			if isvalid {
				num = next - current
				i++
			}
		} else {
			num = val[s[i]]
		}
		res += num
	}
	return res
}

// @leet end

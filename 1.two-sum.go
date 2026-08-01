// @leet start
func twoSum(nums []int, target int) []int {
	for j := range nums {
		for i := j + 1; i < len(nums); i++ {
			if nums[i]+nums[j] == target {
				return []int{j, i}
			}
		}
	}
	return []int{-1, -1}
}

// @leet end

package leetcode

// @leet start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	result := make([]int, 0, len(nums1)+len(nums2))

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	result = append(result, nums1[i:]...)
	result = append(result, nums2[j:]...)

	length := len(result)
	if length%2 == 0 {
		upper := float64(result[length/2])
		lower := float64((result[(length/2)-1]))
		return (upper + lower) / float64(2)
	}
	return float64(result[length/2])
}

// @leet end

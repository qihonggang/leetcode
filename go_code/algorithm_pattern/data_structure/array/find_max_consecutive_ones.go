package array

func findMaxConsecutiveOnes(nums []int) int {
	var l, result int
	for i, v := range nums{
		if v == 1 {
			l++
		}
		if nums[i] != nums[i-1] && v == 0 {
			result = max(l, result)
			l = 0
		}
	}
	result = max(l, result)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

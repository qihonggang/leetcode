package practice

import "testing"

func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(mn*nums[i], nums[i]))
		minF = min(mx*nums[i], min(mn*nums[i], nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestMaxProduct(t *testing.T) {
	nums := []int{5, 30, -3, 4, -3}
	t.Log(maxProduct(nums))
}

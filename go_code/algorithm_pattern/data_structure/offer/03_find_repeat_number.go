package offer

import "sort"

func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	var num int
	for i, v := range nums[1:]{
		if v == nums[i] {
			num = v
			break
		} else {
			num = -1
		}
	}
	return num
}

package array

import "sort"

func findErrorNums(nums []int) []int {
	sameNum := 0
	sum := 0
	temp := 0
	if len(nums) == 0 {
		return nil
	}
	sort.Sort(sort.IntSlice(nums))

	for _, v := range nums {
		if temp == 0 {
			temp = v
		} else {
			if temp == v {
				sameNum = v
			}
			temp = v
		}
		sum += v
	}
	return []int{sameNum, ((1+len(nums)) * len(nums)/2 - (sum-sameNum))}
}

func findErrorNums2(nums []int) []int {
	type void struct {}
	var value void
	set := make(map[int]void)

	sameNum := 0
	length := len(nums)
	sum := 0

	for i := 0; i < length; i++ {
		if _, ok := set[nums[i]]; ok {
			sameNum = nums[i]
		} else {
			set[nums[i]] = value
		}
		sum += nums[i]
	}
	return []int{sameNum, ((1 + length) * length/2) - (sum - sameNum)}
}

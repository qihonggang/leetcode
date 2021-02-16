package main

import (
	"fmt"
)

/**
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

进阶：你可以实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-missing-positive
 */

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			fmt.Println(num - 1)
			nums[num - 1] = - abs(nums[num - 1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
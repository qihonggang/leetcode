package main

/**
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

单调双端队列

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-maximum
 */

func maxSlidingWindow(nums []int, k int) []int {
	// 单调双端队列
	q := []int{}
	
	// 单调双端队列 push 方法
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q) - 1]] {
			q = q[:len(q) - 1]
		}
		q = append(q, i)
	}
	// 0-k
	for i := 0; i < k; i++ {
		push(i)
	}
	// k-n
	n := len(nums)
	ans := make([]int, 1, n - k + 1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i - k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}


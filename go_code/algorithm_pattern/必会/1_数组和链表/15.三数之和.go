package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/**
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
 */

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums) - 2; i++ {
		n1 := nums[i]
		if i > 0 && n1 == nums[i - 1] {
			continue
		}
		l, r := i + 1, len(nums) - 1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1 + n2 + n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1 + n2 + n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	in, _ := reader.ReadString('\n')
	in = strings.Trim(in, "\r\n")
	stringInput := strings.Split(in, ",")
	var nums []int
	for _, v := range stringInput {
		num, _ := strconv.Atoi(v)
		nums = append(nums, num)
	}
	res := threeSum(nums)
	fmt.Print(res)
}
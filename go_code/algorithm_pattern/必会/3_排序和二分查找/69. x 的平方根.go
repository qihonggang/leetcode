package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sqrtx
 */

func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := (r + l) / 2
		if mid * mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	x, _ := strconv.Atoi(input)
	s := mySqrt(x)
	fmt.Print(s)
}
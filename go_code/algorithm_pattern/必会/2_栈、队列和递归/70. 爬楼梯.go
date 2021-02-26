package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

https://leetcode-cn.com/problems/climbing-stairs
 */

func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	n, _ := strconv.Atoi(input)
	mNum := climbStairs(n)
	fmt.Print(mNum)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
题目描述
输入一个int型整数，按照从右向左的阅读顺序，返回一个不含重复数字的新的整数。
保证输入的整数最后一位不是0。
输入描述:
输入一个int型整数

输出描述:
按照从右向左的阅读顺序，返回一个不含重复数字的新的整数
 */
func main() {
	var numm = map[int]bool{}
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.Trim(line, "\r\n")
	nums := []byte(line)
	for i := len(nums) - 1; i >= 0; i-- {
		temp, _ := strconv.Atoi(string(nums[i]))
		if !numm[temp] {
			numm[temp] = true
			fmt.Print(temp)
		}
	}
}
package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
数字颠倒

题目描述
输入一个整数，将这个整数以字符串的形式逆序输出
程序不考虑负数的情况，若数字含有0，则逆序形式也含有0，如输入为100，则输出为001


输入描述:
输入一个int整数

输出描述:
将这个整数以字符串的形式逆序输出
 */
func main() {
	reader := bufio.NewReader(os.Stdin)
	in, _, _ := reader.ReadLine()
	if len(in) == 0 {
		return
	}
	line := string(in)
	for i := len(line) - 1; i >= 0; i -- {
		fmt.Print(line[i:i+1])
	}
}

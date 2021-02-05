package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
题目描述
功能:输入一个正整数，按照从小到大的顺序输出它的所有质因子（重复的也要列举）（如180的质因子为2 2 3 3 5 ）

最后一个数后面也要有空格

输入描述:
输入一个long型整数

输出描述:
按照从小到大的顺序输出它的所有质数的因子，以空格隔开。最后一个数后面也要有空格。
 */
func main() {
	reader := bufio.NewReader(os.Stdin)
	in, _, _ := reader.ReadLine()
	if len(in) == 0 {
		return
	}
	num, _ := strconv.Atoi(string(in))
	for i := 2; i * i <= num; {
		if num % i == 0 {
			fmt.Printf("%d ", i)
			num = num / i
		} else {
			i++
		}
	}
	if num > 1 {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
明明的随机数

题目描述
明明想在学校中请一些同学一起做一项问卷调查，为了实验的客观性，他先用计算机生成了N个1到1000之间的随机整数（N≤1000），对于其中重复的数字，只保留
一个，把其余相同的数去掉，不同的数对应着不同的学生的学号。然后再把这些数从小到大排序，按照排好的顺序去找同学做调查。请你协助明明完成“去重”与“排序
”的工作(同一个测试用例里可能会有多组数据，希望大家能正确处理)。
注：测试用例保证输入参数的正确性，答题者无需验证。测试用例不止一组。
当没有新的输入时，说明输入结束。

输入描述:
注意：输入可能有多组数据。每组数据都包括多行，第一行先输入随机整数的个数N，接下来的N行再输入相应个数的整数。具体格式请看下面的"示例"。
输出描述:
返回多行，处理后的结果
 */

func main() {
	read := bufio.NewReader(os.Stdin)
	for {
		in, _, _ := read.ReadLine()
		if len(in) == 0 {
			break
		}
		count, _ := strconv.Atoi(string(in))
		mark := [1000]bool{}
		for i := 0; i < count; i++ {
			nr, _, _ := read.ReadLine()
			n, _ := strconv.Atoi(string(nr))
			mark[n] = true
		}
		for k, v := range mark {
			if v {
				fmt.Println(k)
			}
		}
	}

}
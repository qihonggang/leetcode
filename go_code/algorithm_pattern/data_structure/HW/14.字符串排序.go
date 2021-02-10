package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

/**
字符串排序

题目描述
给定n个字符串，请对n个字符串按照字典序排列。
输入描述:
输入第一行为一个正整数n(1≤n≤1000),下面n行为n个字符串(字符串长度≤100),字符串中只含有大小写字母。
输出描述:
数据输出n行，输出结果为按照字典序排列的字符串。
 */

func main() {
	reader := bufio.NewReader(os.Stdin)
	in, _, _ := reader.ReadLine()
	if len(in) == 0 {
		return
	}
	num, _ := strconv.Atoi(string(in))
	words := []string{}
	for i := 0; i < num; i++ {
		word, _, _ := reader.ReadLine()
		words = append(words, string(word))
	}
	sort.Strings(words)
	for _, v := range words {
		fmt.Println(v)
	}
}
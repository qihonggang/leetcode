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
题目描述
数据表记录包含表索引和数值（int范围的正整数），请对表索引相同的记录进行合并，即将相同索引的数值进行求和运算，输出按照key值升序进行输出。

输入描述:
先输入键值对的个数
然后输入成对的index和value值，以空格隔开

输出描述:
输出合并后的键值对（多行）
 */

func main() {
	reader := bufio.NewReader(os.Stdin)
	in, _, _ := reader.ReadLine()
	if len(in) == 0 {
		return
	}
	count, _ := strconv.Atoi(string(in))
	kv := make(map[int]int)
	for i:= 0; i < count; i++ {
		line, _, _ := reader.ReadLine()
		lines :=  strings.Split(string(line), " ")
		if len(lines) != 2 {
			continue
		}
		k, _ := strconv.Atoi(lines[0])
		v, _ := strconv.Atoi(lines[1])
		kv[k] += v
	}
	var nums = []int{}
	for k, _ := range kv {
		nums = append(nums, k)
	}
	sort.Ints(nums)
	for _, v := range nums {
		fmt.Println(v, kv[v])
	}
}

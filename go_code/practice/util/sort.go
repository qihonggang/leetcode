package util

import "sort"

func sortUtil() {
	// int排序
	sort.Ints([]int{})

	// 字符串排序
	sort.Strings([]string{})

	// 自定义排序
	s := []int{}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

/**
题目：存在重复元素
*/
package practice

import "testing"

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = 1
		}
	}
	return false
}

func TestContanisDuplicate(t *testing.T) {
	arr := []int{1, 2, 4, 3}
	result := containsDuplicate(arr)
	t.Log(result)
}

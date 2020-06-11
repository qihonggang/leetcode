/**
* 题目：两数之和
 */
package go_code

import (
	"reflect"
	"testing"
)

type IO1 struct {
	in1 []int
	in2 int
	out []int
}

func Test1(t *testing.T) {
	tc := map[string]IO1{
		"1": IO1{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		"2": IO1{[]int{2, 7, 2, 15}, 4, []int{0, 2}},
		"3": IO1{[]int{1, 2, 2, 6}, 4, []int{1, 2}},
	}

	for k, v := range tc {
		out := twoSum1(v.in1, v.in2)
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v:except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}

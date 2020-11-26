package offer

import "testing"

func TestFindRepeatNumber(t *testing.T){
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	num := findRepeatNumber(nums)
	t.Log(num)
}

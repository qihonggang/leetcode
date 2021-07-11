/**
除自身以外数组的乘积
*/
package practice

import "testing"

func productExceptSelf(nums []int) []int {
	var result = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		result[i] = 1
	}

	for i, _ := range nums {
		for j, w := range nums {
			if i != j {
				result[i] *= w
			}
		}
	}
	return result
}

func productExceptSelf2(nums []int) []int {
	length := len(nums)

	L, R, answer := make([]int, length), make([]int, length), make([]int, length)

	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = nums[i-1] * L[i-1]
	}

	R[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}

	for i := 0; i < length; i++ {
		answer[i] = L[i] * R[i]
	}

	return answer
}

func TestExcepteSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf2(nums)
	t.Log(result)
}

package practice
/**
 268. 缺失数字
 */
import "testing"

func missingNumber(nums []int) int {
	res := 0
	for i, n := range nums {
		res ^= i ^ n
	}
	res ^= len(nums)
	return res
}

func TestMissingNuber(t *testing.T) {
	num := []int{3,0,1}
	t.Log(missingNumber(num))
}

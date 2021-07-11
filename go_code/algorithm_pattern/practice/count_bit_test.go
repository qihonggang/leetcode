package practice

/**
338. 比特位计数
*/
import "testing"

func countBits(num int) []int {
	arr := make([]int, num+1)
	for i := 1; i <= num; i++ {
		arr[i] = arr[i>>1] + (i & 1)
	}
	return arr
}

func TestCountBits(t *testing.T) {
	t.Log(countBits(5))
}

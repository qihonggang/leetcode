package practice
/**
191. 位1的个数
*/
import "testing"

func hammingWeight(num uint32) int {
	count := 0
	for 0 < num {
		if num%2 == 1 {
			count++
		}
		num /= 2
	}
	return count
}

func TestHammingWeight(t *testing.T) {
	var num uint32 = 00000000000000000000000000001011
	t.Log(hammingWeight(num))
}

package practice

/**
371. 两整数之和
*/
import "testing"

func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		carry := (a & b) << 1
		a = sum
		b = carry
	}
	return a
}

func TestGetSum(t *testing.T) {
	t.Log(getSum(1, 2))
}

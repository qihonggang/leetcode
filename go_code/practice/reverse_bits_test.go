package practice
/**
190. 颠倒二进制位
 */
import "testing"

func reverseBits(num uint32) uint32 {
	ret := uint32(0)
	power := uint32(31)
	for num != 0 {
		ret += (num & 1) << power
		num = num >> 1
		power -= 1
	}
	return ret
}

func TestReverseBits(t *testing.T){
	t.Log(reverseBits(
		000000101001010))
}

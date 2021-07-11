package practice

/**
190. 颠倒二进制位
*/
import (
	"fmt"
	"testing"
)

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

func TestReverseBits(t *testing.T) {
	num := uint32(964176192)
	fmt.Printf("%b\n", num)
	a := reverseBits(num)
	fmt.Printf("%b\n", a)
}

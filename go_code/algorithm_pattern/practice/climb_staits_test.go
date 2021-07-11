package practice

import "testing"

/*
70. 爬楼梯
*/
func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = q + p
	}
	return r
}

func TestClimbStairs(t *testing.T) {
	t.Log(climbStairs(3))
}

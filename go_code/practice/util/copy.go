package util

func copyUtil() {
	i := 1
	n := 9
	x := 5
	// 删除a[i], 可以copy将i+1到末尾的值覆盖到i,然后末尾-1
	a := make([]int, 0)
	copy(a[i:], a[i+1:])
	a = a[:len(a) - 1]

	//make创建长度，则通过索引复制
	b := make([]int, n)
	b[n] = x

	// make长度为0，则通过append()赋值
	c := make([]int, 0)
	c = append(c, x)
}

package binary_tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 创建树
func TreeCreate(i int, arr []int) *TreeNode{
	t := &TreeNode{arr[i], nil, nil}
	if i < len(arr) && 2 * i + 1 < len(arr) {
		t.Left = TreeCreate(2 * i + 1, arr)
	}
	if i < len(arr) && 2 * i + 2 < len(arr) {
		t.Right = TreeCreate(2 * i + 2, arr)
	}
	return t
}
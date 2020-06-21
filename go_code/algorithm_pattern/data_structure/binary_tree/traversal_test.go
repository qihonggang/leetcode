package binary_tree

import "testing"

func TestTraversal(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	tree := TreeCreate(0, arr)
	// 前序遍历（递归）测试
	preorderTraversal(tree)
	// 前序遍历（非递归）测试
	treeArrPre2 :=  preorderTraversal2(tree)
	t.Logf("序遍历（非递归）：%d", treeArrPre2)
	// 中序遍历（非递归）测试
	treeArrIn := inorderTraversal(tree)
	t.Logf("中序遍历（非递归）：%d", treeArrIn)
	// 后序遍历（非递归）测试
	treeArrPost := postorderTraversal(tree)
	t.Logf("后序遍历（非递归）：%d", treeArrPost)
	// DFS深度搜索-从上到下测试
	treeArrPre3 := preorderTraversal3(tree)
	t.Logf("DFS深度搜索-从上到下：%d", treeArrPre3)
	// DFS深度搜索-从下向上（分治法）测试
	treeArrPre4 := preorderTraversal4(tree)
	t.Logf("DFS深度搜索-从下向上（分治法）：%d", treeArrPre4)
	// BFS层次遍历测试
	treeArrLevel := levelOrder(tree)
	t.Logf("BFS层次遍历：%d", treeArrLevel)
}

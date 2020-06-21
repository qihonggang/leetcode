package binary_tree

import (
	"fmt"
)

// 前序递归
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	// 先访问根再访问左右
	fmt.Print(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// 前序非递归遍历
func preorderTraversal2(root *TreeNode) []int {
	// 非递归
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root !=nil || len(stack) != 0 {
		for root != nil {
			// 前序遍历，所以先保存结果
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		root = node.Right
	}
	return result
}

// 中序非递归遍历
// 思路：通过stack保存已经访问的元素，用于原路返回
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left // 一直向左
		}
		// 弹出
		val := stack[len(stack) - 1]
		stack = stack[:len(stack) -1]
		result = append(result, val.Val)
		root = val.Right
	}
	return result
}

//后序非递归遍历
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	// 通过lastVisit表示右子节点是否已经弹出
	var lastVisit * TreeNode
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 这里先看看，先不弹出
		node := stack[len(stack) - 1]
		// 根节点必须在右节点弹出之后，再弹出
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack) - 1] // pop
			result = append(result, node.Val)
			// 标记当前这个节点已经弹出过
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}

// DFS深度搜索-从上到下
func preorderTraversal3(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}
// v1: 深度白能力，结果指针作为参数传入到函数内部
func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

// DFS深度搜索-从下向上（分治法）
func preorderTraversal4(root *TreeNode) []int {
	result := divideAndConquer(root)
	return result
}
// v2: 通过分治法遍历
func divideAndConquer(root *TreeNode) []int {
	result := make([]int, 0)
	// 返回条件（null & leaf）
	if root == nil {
		return result
	}
	// 分治（Divide）
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	// 合并结果（Conquer）
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// BFS层次遍历
func levelOrder(root *TreeNode) [][]int {
	// 通过上一层的长度确定下一层的元素
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		// 为什么要取length?
		// 记录当前层有多少元素
		l := len(queue)
		for i := 0; i < l; i++ {
			// 出队列
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		result = append(result, list)
	}
	return result
}



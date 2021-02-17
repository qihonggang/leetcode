package main

import "container/heap"

/**
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
*/

// 遍历列表数组
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	var result *ListNode
	for _, list := range lists {
		result = mergeTwoSortedLists(result, list)
	}
	return result
}

func mergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return dummy.Next
}

// 最小堆
func mergeKLists2(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	var h IntHeap
	heap.Init(&h)

	for _, list := range lists {
		if list != nil {
			heap.Push(&h, list)
		}
	}
	dummy := &ListNode{}
	p := dummy

	for h.Len() > 0 {
		min := heap.Pop(&h).(*ListNode)
		p.Next = min
		p = p.Next
		if min.Next != nil {
			heap.Push(&h, min.Next)
		}
	}
	return dummy.Next

}

type IntHeap []*ListNode

func (h IntHeap) Len() int{
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0: n-1]
	return x
}

// 分治算法
func mergeKLists3(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists) - 1)
}

func merge(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	if start > end{
		return nil
	}
	mid := start + (end - start) / 2
	left := merge(lists, start, mid)
	right := merge(lists, mid + 1, end)
	return mergeTwoSortedLists(left, right)
}


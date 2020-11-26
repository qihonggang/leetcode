package offer

import "testing"

func TestReversePrint(t *testing.T) {
	n4 := &ListNode{Val: 4}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	head := &ListNode{Val: 1, Next: n2}
	t.Log(reversePrint(head))
}

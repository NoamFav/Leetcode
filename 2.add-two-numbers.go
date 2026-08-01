package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// @leet start

// Definition for singly-linked list.

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var head *ListNode
	var remainder int
	var digit int
	var val1 int
	var val2 int

	result = &ListNode{}
	head = result

	for l1 != nil || l2 != nil || remainder != 0 {
		val1, val2 = 0, 0
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}
		digit = (val1 + val2 + remainder) % 10
		remainder = (val1 + val2 + remainder) / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		new := &ListNode{Val: digit, Next: nil}
		result.Next = new
		result = result.Next
	}

	return head.Next
}

// @leet end

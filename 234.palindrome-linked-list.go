package leetcode

import "go/types"

type ListNode struct {
	Val  int
	Next *ListNode
}

// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	var num []int
	for ; head != nil; head = head.Next {
		num = append(num, head.Val)
	}

	length := len(num)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		if num[i] != num[j] {
			return false
		}
	}
	return true
}

// @leet end

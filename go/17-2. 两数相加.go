package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	res := head
	k := 0
	for l1 != nil && l2 != nil {
		head.Next = &ListNode{Val: (l1.Val + l2.Val + k) % 10}
		k = (l1.Val + l2.Val + k) / 10
		head = head.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		head.Next = &ListNode{Val: (l1.Val + k) % 10}
		k = (l1.Val + k) / 10
		head = head.Next
		l1 = l1.Next
	}

	for l2 != nil {
		head.Next = &ListNode{Val: (l2.Val + k) % 10}
		k = (l2.Val + k) / 10
		head = head.Next
		l2 = l2.Next
	}

	if k != 0 {
		head.Next = &ListNode{Val: k}
	}

	return res.Next
}

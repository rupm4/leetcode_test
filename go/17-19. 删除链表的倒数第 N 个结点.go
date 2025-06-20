package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{Next: head}
	first, second := head, res
	i := 0
	for first != nil {
		if i < n {
			first = first.Next
		} else {
			first = first.Next
			second = second.Next
		}
		i++
	}

	second.Next = second.Next.Next
	return res.Next
}

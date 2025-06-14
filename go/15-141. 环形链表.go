package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	node1, node2 := head, head
	return isCycle(node1, node2)
}

func isCycle(node1, node2 *ListNode) bool {
	if node1 == nil || node2 == nil || node2.Next == nil {
		return false
	}
	node1 = node1.Next
	node2 = node2.Next.Next

	if node1 == node2 {
		return true
	}

	return isCycle(node1, node2)
}

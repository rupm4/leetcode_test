package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	node1, node2 := head, head
	node3 := isCycle(node1, node2)

	if node3 == nil {
		return nil
	}

	for node3 != head {
		node3 = node3.Next
		head = head.Next
	}

	return node3
}

func isCycle(node1, node2 *ListNode) *ListNode {
	if node1 == nil || node2 == nil || node2.Next == nil {
		return nil
	}
	node1 = node1.Next
	node2 = node2.Next.Next

	if node1 == node2 {
		return node1
	}

	return isCycle(node1, node2)
}

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	lenA := getLenLink(headA)
	lenB := getLenLink(headB)

	l := lenA - lenB

	for l != 0 {
		if l > 0 {
			headA = headA.Next
			l--
		} else {
			headB = headB.Next
			l++
		}
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}

func getLenLink(head *ListNode) int {
	long := 0
	for head != nil {
		long++
		head = head.Next
	}

	return long
}

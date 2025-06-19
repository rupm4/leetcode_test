package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	res := head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			head.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			head.Next = list1
			list1 = list1.Next
		} else {
			if list1.Val >= list2.Val {
				head.Next = list2
				list2 = list2.Next
			} else {
				head.Next = list1
				list1 = list1.Next
			}
		}

		head = head.Next
	}

	return res.Next
}

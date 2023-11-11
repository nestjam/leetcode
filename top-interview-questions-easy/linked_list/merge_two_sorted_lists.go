package linkedlist

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, tail *ListNode

	for list1 != nil || list2 != nil {
		var next *ListNode
		if list1 == nil {
			next = list2
			list2 = nil
		} else if list2 == nil {
			next = list1
			list1 = nil
		} else if list1.Val < list2.Val {
			next = list1
			list1 = list1.Next
		} else {
			next = list2
			list2 = list2.Next
		}

		if head == nil {
			head = next
			tail = next
		} else {
			tail.Next = next
			tail = next
		}
	}

	return head
}

package linkedlist

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

//Constraint
//The number of the nodes in the list is in the range [0, 10^4]
// func hasCycle(head *ListNode) bool {
// 	var i int = 0
// 	for head != nil {
// 		i++
// 		head = head.Next
// 		if i > 10000 {
// 			return true
// 		}
// 	}
// 	return false
// }

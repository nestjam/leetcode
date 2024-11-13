package reverselinkedlistii

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	i := 1

	current := head
	var lastFixed *ListNode
	for i < left {
		lastFixed = current
		current = current.Next
		i++
	}

	reversedTail := current
	reversedHead := current.Next
	next := reversedHead.Next
	reversedHead.Next = current
	current.Next = nil

	for next != nil {
		i++
		if i == right {
			reversedTail.Next = next
			break
		}
		t := next.Next
		next.Next = reversedHead
		reversedHead = next
		next = t
	}

	if lastFixed == nil {
		return reversedHead
	}

	lastFixed.Next = reversedHead
	return head
}

package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{Val: l1.Val + l2.Val, Next: &ListNode{}}
	next := sum.Next
	prev := sum
	l1 = l1.Next
	l2 = l2.Next

	if sum.Val > 9 {
		next.Val += 1
		prev.Val -= 10
	}

	for l1 != nil || l2 != nil {
		prev = next
		next = &ListNode{}
		prev.Next = next

		if l1 != nil {
			prev.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			prev.Val += l2.Val
			l2 = l2.Next
		}

		if prev.Val > 9 {
			next.Val += 1
			prev.Val -= 10
		}
	}

	if next.Val == 0 {
		prev.Next = nil
	}

	return sum
}

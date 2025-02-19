package insertionsortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(unsorted *ListNode) *ListNode {
	sorted, unsorted := unsorted, unsorted.Next
	sorted.Next = nil

	for unsorted != nil {
		node := unsorted
		unsorted = unsorted.Next

		node.Next, sorted = sorted, node

		if sorted.Val <= sorted.Next.Val {
			continue
		}

		cur := sorted
		sorted = sorted.Next
		var prev *ListNode
		for cur.Next != nil && cur.Val > cur.Next.Val {
			next := cur.Next
			if prev != nil {
				prev.Next = next
			}
			cur.Next, next.Next = next.Next, cur
			prev = next
		}
	}

	return sorted
}

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

		cur := sorted
		var prev *ListNode
		for cur != nil && node.Val > cur.Val {
			prev = cur
			cur = cur.Next
		}

		if prev == nil {
			node.Next = cur
			sorted = node
		} else {
			prev.Next = node
			node.Next = cur
		}
	}

	return sorted
}

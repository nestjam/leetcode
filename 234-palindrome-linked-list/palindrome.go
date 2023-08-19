package palindromelinkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var next *ListNode
	var tail *ListNode
	for slow != nil {
		next = tail
		tail = slow
		slow = slow.Next
		tail.Next = next
	}

	for tail != nil {
		if tail.Val != head.Val {
			return false
		}
		tail = tail.Next
		head = head.Next
	}
	return true
}

func getMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	return reversePair(head, head.Next)

	// var next *ListNode
	// var tail *ListNode
	// for head != nil {
	// 	next = tail
	// 	tail = head
	// 	head = head.Next
	// 	tail.Next = next
	// }
	// return tail
}

func reversePair(prev, next *ListNode) *ListNode {
	if next == nil {
		return prev
	}
	tail := reversePair(next, next.Next)
	next.Next = prev
	prev.Next = nil
	return tail
}

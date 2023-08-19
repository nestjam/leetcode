package palindromelinkedlist

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		desc string
		want bool
		list []int
	}{
		{
			desc: "[1]",
			want: true,
			list: []int{1},
		},
		{
			desc: "[1, 2]",
			want: false,
			list: []int{1, 2},
		},
		{
			desc: "[1, 2, 1]",
			want: true,
			list: []int{1, 2, 1},
		},
		{
			desc: "[1, 2, 2, 1]",
			want: true,
			list: []int{1, 2, 2, 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := isPalindrome(newLinkedList(tC.list))

			if got != tC.want {
				t.Errorf("want `%v` but got `%v` with list %v", tC.want, got, tC.list)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		desc string
		want []int
		list []int
	}{
		{
			desc: "[1]",
			want: []int{1},
			list: []int{1},
		},
		{
			desc: "[1, 2]",
			list: []int{1, 2},
			want: []int{2, 1},
		},
		{
			desc: "[1, 2, 3]",
			list: []int{1, 2, 3},
			want: []int{3, 2, 1},
		},
		{
			desc: "[]",
			want: []int{},
			list: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := toList(reverse(newLinkedList(tC.list)))

			if !reflect.DeepEqual(tC.want, got) {
				t.Errorf("want `%v` but got `%v` with list %v", tC.want, got, tC.list)
			}
		})
	}
}

func TestGetMiddle(t *testing.T) {
	testCases := []struct {
		desc string
		want int
		list []int
	}{
		{
			desc: "[1]",
			want: 1,
			list: []int{1},
		},
		{
			desc: "[1, 2]",
			want: 2,
			list: []int{1, 2},
		},
		{
			desc: "[1, 2, 3]",
			want: 2,
			list: []int{1, 2, 3},
		},
		{
			desc: "[1, 2, 3, 4]",
			want: 3,
			list: []int{1, 2, 3, 4},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := getMiddle(newLinkedList(tC.list))

			if !reflect.DeepEqual(tC.want, got.Val) {
				t.Errorf("want `%v` but got `%v` with list %v", tC.want, got.Val, tC.list)
			}
		})
	}
}

func newLinkedList(list []int) *ListNode {
	var head *ListNode
	var next *ListNode

	for i := 0; i < len(list); i++ {
		if i == 0 {
			head = newListNode(list[i])
			next = head
		} else {
			prev := next
			next = newListNode(list[i])
			prev.Next = next
		}
	}
	return head
}

func newListNode(val int) *ListNode {
	return &ListNode{val, nil}
}

func toList(head *ListNode) []int {
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	return list
}

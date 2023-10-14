package linkedlist

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			want: false,
			args: args{
				head: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
		{
			want: true,
			args: args{
				head: newCycle01(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newCycle01() *ListNode {
	var head = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
		},
	}
	head.Next.Next = head
	return head
}

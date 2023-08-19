package minimumdistancebetweenbstnodes

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		tree *TreeNode
		want int
	}{
		{
			desc: "4,1,null -> 3",
			tree: newNode(4, newLeaveNode(1), nil),
			want: 3,
		},
		{
			desc: "3,1,4 -> 1",
			tree: newNode(3, newLeaveNode(1), newLeaveNode(4)),
			want: 1,
		},
		{
			desc: "4,2,6,1,3 -> 1",
			tree: newNode(4, newNode(2, newLeaveNode(1), newLeaveNode(3)), newLeaveNode(6)),
			want: 1,
		},
		{
			desc: "1,0,48,null,null,12,49 -> 1",
			tree: newNode(1, newLeaveNode(0), newNode(48, newLeaveNode(12), newLeaveNode(49))),
			want: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := minDiffInBST(tC.tree)

			if got != tC.want {
				t.Errorf("expected %v but got %v", tC.want, got)
			}
		})
	}
}

func newNode(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

func newLeaveNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

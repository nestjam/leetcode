package binarytreepreordertraversal

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	testCases := []struct {
		desc string
		tree *TreeNode
		want []int
	}{
		{
			desc: "1",
			tree: &TreeNode{Val: 1},
			want: []int{1},
		},
		{
			desc: "",
			tree: nil,
			want: []int { },
		},
		{
			desc: "{1 l2 r3}",
			tree: &TreeNode{ Val: 1, Left: &TreeNode{ Val: 2 }, Right: &TreeNode{ Val: 3 } },
			want: []int { 1, 2, 3 },
		},
		{
			desc: "{1 l2 r3 2l:4}",
			tree: &TreeNode{ Val: 1, Left: &TreeNode{ Val: 2, Left: &TreeNode{ Val: 4 } }, Right: &TreeNode{ Val: 3 } },
			want: []int { 1, 2, 4, 3 },
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := preorderTraversal(tC.tree)

			if !reflect.DeepEqual(tC.want, got) {
				t.Errorf("want %v but got %v", tC.want, got)
			}
		})
	}
}

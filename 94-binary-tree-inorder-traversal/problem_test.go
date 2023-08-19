package binarytreeinordertraversal

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	testCases := []struct {
		desc string
		tree *TreeNode
		want []int
	}{
		{
			desc: "1",
			tree: &TreeNode{ Val: 1 },
			want: []int { 1 },
		},
		{
			desc: "",
			tree: nil,
			want: []int { },
		},
		{
			desc: "{1 l2 r3}",
			tree: &TreeNode{ Val: 1, Left: &TreeNode{ Val: 2 }, Right: &TreeNode{ Val: 3 } },
			want: []int { 2, 1, 3 },
		},
		{
			desc: "{1 l1.2 r1.3 l2.4}",
			tree: &TreeNode{ Val: 1, Left: &TreeNode{ Val: 2, Left: &TreeNode{ Val: 4 } }, Right: &TreeNode{ Val: 3 } },
			want: []int { 4, 2, 1, 3 },
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := inorderTraversal(tC.tree)

			if !reflect.DeepEqual(tC.want, got) {
				t.Errorf("want %v but got %v", tC.want, got)
			}
		})
	}
}

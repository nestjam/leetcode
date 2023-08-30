package trees

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		root *TreeNode
		want int
		desc string
	}{
		{
			root: nil,
			want: 0,
		},
		{
			root: &TreeNode{},
			want: 1,
		},
		{
			root: &TreeNode{
				Left: &TreeNode{},
			},
			want: 2,
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d %v", i + 1, tC.desc), func(t *testing.T) {
			got:= maxDepth(tC.root)

			if got != tC.want {
				t.Errorf("expected %v but %v", tC.want, got)
			}
		})
	}
}
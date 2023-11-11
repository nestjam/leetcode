package trees

import (
	"fmt"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	testCases := []struct {
		root *TreeNode
		want bool
		desc string
	}{
		{
			root: &TreeNode{},
			want: true,
		},
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: false,
		},
		{
			root: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
				},
			},
			want: true,
		},
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
			},
			want: true,
		},
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 10,
					},
				},
			},
			want: false,
		},
		{
			root: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 0,
					},
				},
			},
			want: false,
		},
		{
			root: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 0,
					},
				},
			},
			want: false,
		},
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			want: false,
		},
		{
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			want: false,
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d %v", i+1, tC.desc), func(t *testing.T) {
			got := isValidBST(tC.root)

			if got != tC.want {
				t.Errorf("expected %v but %v", tC.want, got)
			}
		})
	}
}

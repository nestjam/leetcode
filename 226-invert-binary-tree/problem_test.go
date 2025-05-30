package invertbinarytree

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				root: nil,
			},
			want: nil,
		},
		{
			args: args{
				root: &TreeNode{Val: 1},
			},
			want: &TreeNode{Val: 1},
		},
		{
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2},
				},
			},
			want: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
		},
		{
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Right: &TreeNode{Val: 3},
					},
				},
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
		},
		{
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val:   2,
						Right: &TreeNode{Val: 3},
					},
				},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

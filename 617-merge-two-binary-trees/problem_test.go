package mergetwobinarytrees

import (
	"reflect"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				root1: nil,
				root2: nil,
			},
			want: nil,
		},
		{
			args: args{
				root1: &TreeNode{Val: 1},
				root2: nil,
			},
			want: &TreeNode{Val: 1},
		},
		{
			args: args{
				root2: &TreeNode{Val: 1},
				root1: nil,
			},
			want: &TreeNode{Val: 1},
		},
		{
			args: args{
				root2: &TreeNode{Val: 1},
				root1: &TreeNode{Val: 2},
			},
			want: &TreeNode{Val: 3},
		},
		{
			args: args{
				root1: &TreeNode{Val: 1},
				root2: &TreeNode{Val: 2, Left: &TreeNode{}},
			},
			want: &TreeNode{Val: 3, Left: &TreeNode{}},
		},
		{
			args: args{
				root1: &TreeNode{Val: 1},
				root2: &TreeNode{Val: 2, Right: &TreeNode{}},
			},
			want: &TreeNode{Val: 3, Right: &TreeNode{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

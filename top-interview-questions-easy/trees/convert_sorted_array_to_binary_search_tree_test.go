package trees

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				nums: []int{0},
			},
			want: &TreeNode{
				Val: 0,
			},
		},
		{
			args: args{
				nums: []int{-1, 0},
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -1,
				},
			},
		},
		{
			args: args{
				nums: []int{-1, 0, 1},
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -1,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		{
			args: args{
				nums: []int{0, 1, 2, 3},
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		{
			args: args{
				nums: []int{0, 1, 2, 3, 4},
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedArrayToBST(tt.args.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

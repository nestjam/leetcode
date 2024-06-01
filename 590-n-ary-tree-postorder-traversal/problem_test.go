package narytreepostordertraversal

import (
	"reflect"
	"testing"
)

func Test_postorder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				root: nil,
			},
			want: []int{},
		},
		{
			args: args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						{
							Val: 2,
						},
					},
				},
			},
			want: []int{2, 1},
		},
		{
			args: args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						{
							Val: 2,
							Children: []*Node{
								{
									Val: 3,
								},
							},
						},
						{
							Val: 4,
						},
					},
				},
			},
			want: []int{3, 2, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

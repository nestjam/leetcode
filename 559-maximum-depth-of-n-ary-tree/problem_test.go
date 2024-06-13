package maximumdepthofnarytree

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				root: &Node{},
			},
			want: 1,
		},
		{
			args: args{
				root: nil,
			},
			want: 0,
		},
		{
			args: args{
				root: &Node{
					Children: []*Node{
						{},
					},
				},
			},
			want: 2,
		},
		{
			args: args{
				root: &Node{
					Children: []*Node{
						{},
						{
							Children: []*Node{
								{},
							},
						},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

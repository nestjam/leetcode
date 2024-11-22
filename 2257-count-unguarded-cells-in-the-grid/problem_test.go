package countunguardedcellsinthegrid

import "testing"

func Test_countUnguarded(t *testing.T) {
	type args struct {
		m      int
		n      int
		guards [][]int
		walls  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				m:      4,
				n:      6,
				guards: [][]int{{0, 0}, {1, 1}, {2, 3}},
				walls:  [][]int{{0, 1}, {2, 2}, {1, 4}},
			},
			want: 7,
		},
		{
			args: args{
				m:      3,
				n:      3,
				guards: [][]int{{1, 1}},
				walls:  [][]int{{1, 0}, {2, 1}, {1, 2}},
			},
			want: 4,
		},
		{
			args: args{
				m:      3,
				n:      1,
				guards: [][]int{{1, 0}},
				walls:  [][]int{{0, 0}},
			},
			want: 0,
		},
		{
			args: args{
				m:      3,
				n:      1,
				guards: [][]int{{0, 0}},
				walls:  [][]int{{1, 0}},
			},
			want: 1,
		},
		{
			args: args{
				m:      3,
				n:      1,
				guards: [][]int{{2, 0}},
				walls:  [][]int{{1, 0}},
			},
			want: 1,
		},
		{
			args: args{
				m:      3,
				n:      1,
				guards: [][]int{{1, 0}},
				walls:  [][]int{{2, 0}},
			},
			want: 0,
		},
		{
			args: args{
				m:      1,
				n:      3,
				guards: [][]int{{0, 2}},
				walls:  [][]int{{0, 1}},
			},
			want: 1,
		},
		{
			args: args{
				m:      1,
				n:      3,
				guards: [][]int{{0, 1}},
				walls:  [][]int{{0, 2}},
			},
			want: 0,
		},
		{
			args: args{
				m:      1,
				n:      3,
				guards: [][]int{{0, 0}},
				walls:  [][]int{{0, 1}},
			},
			want: 1,
		},
		{
			args: args{
				m:      1,
				n:      1,
				guards: [][]int{{0, 0}},
				walls:  [][]int{},
			},
			want: 0,
		},
		{
			args: args{
				m:      1,
				n:      2,
				guards: [][]int{},
				walls:  [][]int{{0, 0}},
			},
			want: 1,
		},
		{
			args: args{
				m:      1,
				n:      2,
				guards: [][]int{},
				walls:  [][]int{},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUnguarded(tt.args.m, tt.args.n, tt.args.guards, tt.args.walls); got != tt.want {
				t.Errorf("countUnguarded() = %v, want %v", got, tt.want)
			}
		})
	}
}

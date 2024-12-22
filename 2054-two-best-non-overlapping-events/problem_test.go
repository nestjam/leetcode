package twobestnonoverlappingevents

import "testing"

func Test_maxTwoEvents(t *testing.T) {
	type args struct {
		events [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				events: [][]int{{0, 1, 1}, {0, 2, 2}},
			},
			want: 2,
		},
		{
			args: args{
				events: [][]int{{0, 1, 1}, {2, 3, 1}},
			},
			want: 2,
		},
		{
			args: args{
				events: [][]int{{0, 1, 1}, {2, 3, 1}, {0, 3, 3}},
			},
			want: 3,
		},
		{
			args: args{
				events: [][]int{{0, 1, 1}, {2, 3, 1}, {0, 3, 1}},
			},
			want: 2,
		},
		{
			args: args{
				events: [][]int{{0, 1, 1}, {0, 1, 2}, {2, 3, 1}},
			},
			want: 3,
		},
		{
			args: args{
				events: [][]int{{10, 83, 53}, {63, 87, 45}, {97, 100, 32}, {51, 61, 16}},
			},
			want: 85,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTwoEvents(tt.args.events); got != tt.want {
				t.Errorf("maxTwoEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

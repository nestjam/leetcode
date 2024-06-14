package minimumtimevisitingallpoints

import "testing"

func Test_minTimeToVisitAllPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				points: [][]int{{0, 0}},
			},
			want: 0,
		},
		{
			args: args{
				points: [][]int{{0, 0}, {1, 0}},
			},
			want: 1,
		},
		{
			args: args{
				points: [][]int{{0, 0}, {-1, 0}},
			},
			want: 1,
		},
		{
			args: args{
				points: [][]int{{0, 0}, {0, 1}},
			},
			want: 1,
		},
		{
			args: args{
				points: [][]int{{0, 0}, {0, -1}},
			},
			want: 1,
		},
		{
			args: args{
				points: [][]int{{0, 0}, {1, 1}},
			},
			want: 1,
		},
		{
			args: args{
				points: [][]int{{0, 0}, {0, 1}, {1, 1}, {1, 0}, {0, 0}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTimeToVisitAllPoints(tt.args.points); got != tt.want {
				t.Errorf("minTimeToVisitAllPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

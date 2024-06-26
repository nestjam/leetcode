package bookingmeetingroom

import "testing"

func Test_getMaxBookings(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				intervals: [][]int{{1, 2}},
			},
			want: 1,
		},
		{
			args: args{
				intervals: [][]int{{1, 2}, {3, 4}},
			},
			want: 2,
		},
		{
			args: args{
				intervals: [][]int{{1, 2}, {1, 3}},
			},
			want: 1,
		},
		{
			args: args{
				intervals: [][]int{{1, 3}, {2, 3}, {4, 5}},
			},
			want: 2,
		},
		{
			args: args{
				intervals: [][]int{{1, 2}, {2, 3}, {4, 5}, {4, 5}, {5, 6}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxBookings(tt.args.intervals); got != tt.want {
				t.Errorf("getMaxBookings() = %v, want %v", got, tt.want)
			}
		})
	}
}

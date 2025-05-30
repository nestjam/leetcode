package averagewaitingtime

import "testing"

func Test_averageWaitingTime(t *testing.T) {
	type args struct {
		customers [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				customers: [][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}},
			},
			want: 3.25,
		},
		{
			args: args{
				customers: [][]int{{1, 2}, {2, 5}, {4, 3}},
			},
			want: 5,
		},
		{
			args: args{
				customers: [][]int{{0, 2}, {1, 1}},
			},
			want: 2,
		},
		{
			args: args{
				customers: [][]int{{0, 1}, {2, 1}},
			},
			want: 1,
		},
		{
			args: args{
				customers: [][]int{{0, 2}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageWaitingTime(tt.args.customers); got != tt.want {
				t.Errorf("averageWaitingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

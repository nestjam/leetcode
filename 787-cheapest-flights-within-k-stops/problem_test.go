package cheapestflightswithinkstops

import "testing"

func Test_findCheapestPrice(t *testing.T) {
	type args struct {
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:       2,
				flights: [][]int{{0, 1, 100}},
				src:     0,
				dst:     0,
				k:       1,
			},
			want: 0,
		},
		{
			args: args{
				n:       2,
				flights: [][]int{{0, 1, 100}},
				src:     0,
				dst:     1,
				k:       1,
			},
			want: 100,
		},
		{
			args: args{
				n:       2,
				flights: [][]int{{0, 1, 100}},
				src:     0,
				dst:     2,
				k:       1,
			},
			want: -1,
		},
		{
			args: args{
				n:       2,
				flights: [][]int{{0, 1, 100}, {1, 0, 100}},
				src:     0,
				dst:     1,
				k:       1,
			},
			want: 100,
		},
		{
			args: args{
				n:       3,
				flights: [][]int{{0, 1, 100}, {1, 2, 100}},
				src:     0,
				dst:     2,
				k:       1,
			},
			want: 200,
		},
		{
			args: args{
				n:       2,
				flights: [][]int{{0, 1, 100}, {1, 0, 100}},
				src:     0,
				dst:     2,
				k:       1,
			},
			want: -1,
		},
		{
			args: args{
				n:       3,
				flights: [][]int{{0, 1, 100}, {1, 2, 100}, {2, 3, 100}},
				src:     0,
				dst:     3,
				k:       1,
			},
			want: -1,
		},
		{
			args: args{
				n:       4,
				flights: [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}},
				src:     0,
				dst:     3,
				k:       1,
			},
			want: 700,
		},
		{
			args: args{
				n:       3,
				flights: [][]int{{0, 1, 300}, {0, 2, 100}, {2, 1, 100}},
				src:     0,
				dst:     1,
				k:       0,
			},
			want: 300,
		},
		{
			args: args{
				n:       3,
				flights: [][]int{{0, 1, 300}, {0, 2, 100}, {2, 1, 100}},
				src:     0,
				dst:     1,
				k:       1,
			},
			want: 200,
		},
		{
			args: args{
				n:       3,
				flights: [][]int{{0, 1, 100}, {0, 2, 100}, {1, 3, 100}, {2, 3, 100}},
				src:     0,
				dst:     3,
				k:       0,
			},
			want: -1,
		},
		{
			args: args{
				n:       4,
				flights: [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}},
				src:     0,
				dst:     3,
				k:       1,
			},
			want: 6,
		},
		{
			args: args{
				n:       5,
				flights: [][]int{{0, 1, 5}, {0, 3, 2}, {1, 2, 5}, {3, 1, 2}, {1, 4, 1}, {4, 2, 1}},
				src:     0,
				dst:     2,
				k:       2,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCheapestPrice(tt.args.n, tt.args.flights, tt.args.src, tt.args.dst, tt.args.k); got != tt.want {
				t.Errorf("findCheapestPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

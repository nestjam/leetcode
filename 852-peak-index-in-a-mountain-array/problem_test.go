package peakindexinamountainarray

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				arr: []int{0, 1, 0},
			},
			want: 1,
		},
		{
			args: args{
				arr: []int{0, 2, 1, 0},
			},
			want: 1,
		},
		{
			args: args{
				arr: []int{0, 1, 2, 0},
			},
			want: 2,
		},
		{
			args: args{
				arr: []int{0, 1, 2, 3, 2, 1, 0},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

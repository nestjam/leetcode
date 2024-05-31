package arraypartition

import "testing"

func Test_arrayPairSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{2, 1},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{2, 3},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{3, 2, 1, 0},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{1, 4, 3, 2},
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{6, 2, 6, 5, 1, 2},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayPairSum(tt.args.nums); got != tt.want {
				t.Errorf("arrayPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

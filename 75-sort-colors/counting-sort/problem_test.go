package countingsort

import "testing"

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{1, 0, 2},
			},
			want: []int{0, 1, 2},
		},
		{
			args: args{
				nums: []int{0, 1, 1},
			},
			want: []int{0, 1, 1},
		},
		{
			args: args{
				nums: []int{0, 2, 1, 2, 1},
			},
			want: []int{0, 1, 1, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
		})
	}
}

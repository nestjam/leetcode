package numberofgoodpairs

import "testing"

func Test_numIdenticalPairs(t *testing.T) {
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
				nums: []int{1, 1},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{1, 2},
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{1, 2, 1},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{1, 2, 3, 1, 1, 3},
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{1, 1, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIdenticalPairs(tt.args.nums); got != tt.want {
				t.Errorf("numIdenticalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

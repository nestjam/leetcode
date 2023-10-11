package other

import "testing"

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			want: 1,
			args: args{
				nums: []int{0, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

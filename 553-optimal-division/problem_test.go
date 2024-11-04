package optimaldivision

import "testing"

func Test_optimalDivision(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				nums: []int{2},
			},
			want: "2",
		},
		{
			args: args{
				nums: []int{10, 2},
			},
			want: "10/2",
		},
		{
			args: args{
				nums: []int{100, 10, 2},
			},
			want: "100/(10/2)",
		},
		{
			args: args{
				nums: []int{2, 3, 4},
			},
			want: "2/(3/4)",
		},
		{
			args: args{
				nums: []int{1000, 100, 10, 2},
			},
			want: "1000/(100/10/2)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optimalDivision(tt.args.nums); got != tt.want {
				t.Errorf("optimalDivision() = %v, want %v", got, tt.want)
			}
		})
	}
}

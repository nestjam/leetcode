package lemonadechange

import "testing"

func Test_lemonadeChange(t *testing.T) {
	type args struct {
		bills []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				bills: []int{5},
			},
			want: true,
		},
		{
			args: args{
				bills: []int{10},
			},
			want: false,
		},
		{
			args: args{
				bills: []int{20},
			},
			want: false,
		},
		{
			args: args{
				bills: []int{5, 20},
			},
			want: false,
		},
		{
			args: args{
				bills: []int{5, 5, 10, 20},
			},
			want: true,
		},
		{
			args: args{
				bills: []int{5, 5, 5, 20},
			},
			want: true,
		},
		{
			args: args{
				bills: []int{5, 10, 10},
			},
			want: false,
		},
		{
			args: args{
				bills: []int{5, 10, 5, 10},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lemonadeChange(tt.args.bills); got != tt.want {
				t.Errorf("lemonadeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

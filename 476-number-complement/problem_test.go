package numbercomplement

import "testing"

func Test_findComplement(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				num: 0,
			},
			want: 1,
		},
		{
			args: args{
				num: 5,
			},
			want: 2,
		},
		{
			args: args{
				num: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findComplement(tt.args.num); got != tt.want {
				t.Errorf("findComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}

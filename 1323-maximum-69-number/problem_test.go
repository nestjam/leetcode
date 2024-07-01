package maximum69number

import "testing"

func Test_maximum69Number(t *testing.T) {
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
				num: 6,
			},
			want: 9,
		},
		{
			args: args{
				num: 9,
			},
			want: 9,
		},
		{
			args: args{
				num: 69,
			},
			want: 99,
		},
		{
			args: args{
				num: 99,
			},
			want: 99,
		},
		{
			args: args{
				num: 9996,
			},
			want: 9999,
		},
		{
			args: args{
				num: 9669,
			},
			want: 9969,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum69Number(tt.args.num); got != tt.want {
				t.Errorf("maximum69Number() = %v, want %v", got, tt.want)
			}
		})
	}
}

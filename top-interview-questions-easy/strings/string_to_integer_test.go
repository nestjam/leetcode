package strings

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			want: 42,
			args: args{
				s: "42",
			},
		},
		{
			want: 1234567890,
			args: args{
				s: "1234567890",
			},
		},
		{
			want: 123,
			args: args{
				s: "00123",
			},
		},
		{
			want: -123,
			args: args{
				s: "-123",
			},
		},
		{
			want: 1,
			args: args{
				s: "0000000000000000000000000000000000000000000001",
			},
		},
		{
			want: -123,
			args: args{
				s: "  -123",
			},
		},
		{
			want: -2147483648,
			args: args{
				s: "-2147483649",
			},
		},
		{
			want: 2147483647,
			args: args{
				s: "2147483648",
			},
		},
		{
			want: 12,
			args: args{
				s: "12 fa 11",
			},
		},
		{
			want: 12,
			args: args{
				s: "12 -11",
			},
		},
		{
			want: 1,
			args: args{
				s: "+1",
			},
		},
		{
			want: 0,
			args: args{
				s: "+-1",
			},
		},
		{
			want: 0,
			args: args{
				s: "0 1",
			},
		},
		{
			want: 2147483647,
			args: args{
				s: "18446744073709551617",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

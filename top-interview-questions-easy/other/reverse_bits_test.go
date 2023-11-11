package other

import (
	"math/bits"
	"testing"
)

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			args: args{
				num: 0,
			},
			want: 0,
		},
		{
			args: args{
				num: 1,
			},
			want: bits.Reverse32(1),
		},
		{
			args: args{
				num: 43261596,
			},
			want: 964176192,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

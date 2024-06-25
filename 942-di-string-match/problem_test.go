package distringmatch

import (
	"reflect"
	"testing"
)

func Test_diStringMatch(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				s: "D",
			},
			want: []int{1, 0},
		},
		{
			args: args{
				s: "I",
			},
			want: []int{0, 1},
		},
		{
			args: args{
				s: "DD",
			},
			want: []int{2, 1, 0},
		},
		{
			args: args{
				s: "DI",
			},
			want: []int{2, 0, 1},
		},
		{
			args: args{
				s: "DDI",
			},
			want: []int{3, 2, 0, 1},
		},
		{
			args: args{
				s: "IDID",
			},
			want: []int{0, 4, 1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diStringMatch(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

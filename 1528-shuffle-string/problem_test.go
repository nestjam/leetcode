package shufflestring

import "testing"

func Test_restoreString(t *testing.T) {
	type args struct {
		s       string
		indices []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s:       "abc",
				indices: []int{0, 1, 2},
			},
			want: "abc",
		},
		{
			args: args{
				s:       "cba",
				indices: []int{2, 1, 0},
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreString(tt.args.s, tt.args.indices); got != tt.want {
				t.Errorf("restoreString() = %v, want %v", got, tt.want)
			}
		})
	}
}

package jewelsandstones

import "testing"

func Test_count(t *testing.T) {
	type args struct {
		stones string
		jewels string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				stones: "aAAbbbb",
				jewels: "aA",
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				stones: "ZZ",
				jewels: "z",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStones(tt.args.jewels, tt.args.stones); got != tt.want {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}

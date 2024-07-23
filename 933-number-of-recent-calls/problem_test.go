package numberofrecentcalls

import "testing"

func TestRecentCounter_Ping(t *testing.T) {
	type args struct {
		calls []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				calls: []int{},
			},
			want: 0,
		},
		{
			args: args{
				calls: []int{1},
			},
			want: 1,
		},
		{
			args: args{
				calls: []int{1, 2},
			},
			want: 2,
		},
		{
			args: args{
				calls: []int{1, 2, 3002},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := Constructor()
			var got int
			for i := 0; i < len(tt.args.calls); i++ {
				got = counter.Ping(tt.args.calls[i])
			}

			if got != tt.want {
				t.Errorf("RecentCounter.Ping() = %v, want %v", got, tt.want)
			}
		})
	}
}

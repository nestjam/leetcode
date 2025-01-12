package mapsumpairs

import "testing"

func TestMapSum_Sum(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name string
		this *MapSum
		args args
		want int
	}{
		{
			this: func() *MapSum {
				m := Constructor()
				m.Insert("a", 1)
				return &m
			}(),
			args: args{
				prefix: "a",
			},
			want: 1,
		},
		{
			this: func() *MapSum {
				m := Constructor()
				m.Insert("a", 1)
				m.Insert("ab", 1)
				return &m
			}(),
			args: args{
				prefix: "a",
			},
			want: 2,
		},
		{
			this: func() *MapSum {
				m := Constructor()
				m.Insert("apple", 3)
				m.Insert("app", 2)
				return &m
			}(),
			args: args{
				prefix: "ap",
			},
			want: 5,
		},
		{
			this: func() *MapSum {
				m := Constructor()
				m.Insert("apple", 3)
				m.Insert("app", 2)
				m.Insert("apple", 1)
				return &m
			}(),
			args: args{
				prefix: "ap",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Sum(tt.args.prefix); got != tt.want {
				t.Errorf("MapSum.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

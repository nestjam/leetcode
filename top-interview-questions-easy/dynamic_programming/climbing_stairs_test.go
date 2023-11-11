package dynamicprogramming

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	testCases := []struct {
		n    int
		want int
		desc string
	}{
		{
			n:    1,
			want: 1,
		},
		{
			n:    2,
			want: 2,
		},
		{
			n:    3,
			want: 3,
		},
		{
			n:    4,
			want: 5,
		},
		{
			n:    45,
			want: 1836311903,
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d %v", i+1, tC.desc), func(t *testing.T) {
			got := climbStairs(tC.n)

			if tC.want != got {
				t.Errorf("expected %v but %v", tC.want, got)
			}
		})
	}
}

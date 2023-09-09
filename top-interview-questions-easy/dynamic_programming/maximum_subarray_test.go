package dynamicprogramming

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	testCases := []struct {
		array []int
		want int
		desc string
	}{
		{
			array: []int{ 1 },
			want: 1,
		},
		{
			array: []int{ 2, -1 },
			want: 2,
			desc: "[2]",
		},
		{
			array: []int{ 1, 1 },
			want: 2,
			desc: "[1, 1]",
		},
		{
			array: []int{ -1, 1 },
			want: 1,
			desc: "[1]",
		},
		{
			array: []int{ 2, -1, 2 },
			want: 3,
			desc: "[2, -1, 2]",
		},
		{
			array: []int{ -2, -1 },
			want: -1,
			desc: "[-1]",
		},
		{
			array: []int{ -1, -2 },
			want: -1,
			desc: "[-1]",
		},
		{
			array: []int{ -2,1,-3,4,-1,2,1,-5,4 },
			want: 6,
			desc: "[4,-1,2,1]",
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d %v", i + 1, tC.desc), func(t *testing.T) {
			got := maxSubArray(tC.array)

			if tC.want != got {
				t.Errorf("expected %v but %v", tC.want, got)
			}
		})
	}
}
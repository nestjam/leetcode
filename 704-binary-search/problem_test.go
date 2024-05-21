package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		array  []int
		target int
		want   int
		desc   string
	}{
		{
			array:  []int{},
			target: 1,
			want:   -1,
			desc:   "empty",
		},
		{
			array:  []int{1},
			target: 1,
			want:   0,
		},
		{
			array:  []int{2},
			target: 1,
			want:   -1,
			desc:   "not found",
		},
		{
			array:  []int{-1, 0, 1, 2},
			target: 0,
			want:   1,
		},
		{
			array:  []int{-1, 0, 1, 2},
			target: 2,
			want:   3,
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d. %v", i+1, tC.desc), func(t *testing.T) {
			got := search(tC.array, tC.target)

			if got != tC.want {
				t.Errorf("expected %v but %v", tC.want, got)
			}
		})
	}
}

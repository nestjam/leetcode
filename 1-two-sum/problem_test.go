package twosum

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		nums   []int
		target int
		want []int
	}{
		{
			desc: "1",
			nums:   []int{1, 2},
			target: 3,
			want: []int{0, 1},
		},
		{
			desc: "2",
			nums:   []int{1, 2, 3},
			target: 4,
			want: []int{0, 2},
		},
		{
			desc: "3",
			nums:   []int{2,7,11,15},
			target: 9,
			want: []int{0, 1},
		},
		{
			desc: "4",
			nums:   []int{3, 3},
			target: 6,
			want: []int{0, 1},
		},
		{
			desc: "5",
			nums:   []int{-3, 3},
			target: 0,
			want: []int{0, 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := twoSum(tC.nums, tC.target)

			if !reflect.DeepEqual(tC.want, got) {
				t.Errorf("expected %v but got %v", tC.want, got)
			}
		})
	}
}

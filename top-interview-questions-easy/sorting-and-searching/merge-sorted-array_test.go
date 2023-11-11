package sortingandsearching

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	testCases := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
		desc  string
	}{
		{
			nums1: []int{1, 0},
			m:     1,
			nums2: []int{2},
			n:     1,
			want:  []int{1, 2},
		},
		{
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1},
		},
		{
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			want:  []int{1},
		},
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d %v", i+1, tC.desc), func(t *testing.T) {
			merge(tC.nums1, tC.m, tC.nums2, tC.n)

			if !reflect.DeepEqual(tC.want, tC.nums1) {
				t.Errorf("expected %v but %v", tC.want, tC.nums1)
			}
		})
	}
}

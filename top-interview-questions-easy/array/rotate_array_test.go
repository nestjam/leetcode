package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		array []int
		k int
		want []int
		desc string
	}{
		{
			array: []int{ 1, 2 },
			k: 0,
			want: []int{ 1, 2 },
		},
		{
			array: []int{ 1, 2 },
			k: 1,
			want: []int{ 2, 1 },
		},
		{
			array: []int{ 1,2,3,4,5,6,7 },
			k: 3,
			want: []int{ 5,6,7,1,2,3,4 },
		},
		{
			array: []int{ -1 },
			k: 2,
			want: []int{ -1 },
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d %v", i + 1, tC.desc), func(t *testing.T) {
			rotate(tC.array, tC.k)

			if !reflect.DeepEqual(tC.want, tC.array) {
				t.Errorf("expected %v but %v", tC.want, tC.array)
			}
		})
	}
}
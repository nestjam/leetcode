package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		array []int
		want []int
		desc string
	}{
		{
			array: []int{ 1 },
			want: []int{ 1 },
		},
		{
			array: []int{},
			want: []int{},
			desc: "empty",
		},
		{
			array: []int{ 1, 1, 2 },
			want: []int{ 1, 2 },
		},
		{
			array: []int{ 0,0,1,1,1,2,2,3,3,4 },
			want: []int{ 0,1,2,3,4 },
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d %v", i + 1, tC.desc), func(t *testing.T) {
			k := removeDuplicates(tC.array)
			got := tC.array[0:k]

			if !reflect.DeepEqual(tC.want, got) {
				t.Errorf("expected %v but %v", tC.want, got)
			}
		})
	}
}
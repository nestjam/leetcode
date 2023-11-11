package other

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		numRows int
		want    [][]int
		desc    string
	}{
		{
			numRows: 1,
			want:    [][]int{{1}},
		},
		{
			numRows: 2,
			want:    [][]int{{1}, {1, 1}},
		},
		{
			numRows: 3,
			want:    [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			numRows: 5,
			want:    [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d %v", i+1, tC.desc), func(t *testing.T) {
			got := generate(tC.numRows)

			if !reflect.DeepEqual(tC.want, got) {
				t.Errorf("expected %v but %v", tC.want, got)
			}
		})
	}
}

package strings

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		x int
		want int
		desc string
	}{
		{
			x: 1,
			want: 1,
		},
		{
			x: 12,
			want: 21,
		},
		{
			x: 10,
			want: 1,
		},
		{
			x: -12,
			want: -21,
			desc: "negative int",
		},
		{
			x: 2147483647,
			want: 0,
			desc: "max int32",
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d %v", i + 1, tC.desc), func(t *testing.T) {
			got := reverse(tC.x)

			if !reflect.DeepEqual(tC.want, got) {
				t.Errorf("expected %v but %v", tC.want, got)
			}
		})
	}
}
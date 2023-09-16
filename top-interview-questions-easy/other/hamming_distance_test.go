package other

import (
	"fmt"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	testCases := []struct {
		x int
		y int
		want  int
		name  string
	}{
		{
			x: 1,
			y: 1,
			want:  0,
		},
		{
			x: 1,
			y: 0,
			want:  1,
		},
		{
			x: 1,
			y: 4,
			want:  2,
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d %v", i+1, tC.name), func(t *testing.T) {
			got := hammingDistance(tC.x, tC.y)

			if tC.want != got {
				t.Errorf("expected %v but %v", tC.want, got)
			}
		})
	}
}
package sortingandsearching

import (
	"fmt"
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	testCases := []struct {
		versions states
		n        int
		want     int
		desc     string
	}{
		{
			versions: []bool{false},
			n:        1,
			want:     1,
			desc:     "",
		},
		{
			versions: []bool{false, true},
			n:        2,
			want:     2,
			desc:     "",
		},
		{
			versions: []bool{true, true},
			n:        2,
			want:     1,
			desc:     "",
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d %v", i+1, tC.desc), func(t *testing.T) {
			got := firstBadVersion(tC.n, tC.versions)

			if tC.want != got {
				t.Errorf("expected %v but %v", tC.want, got)
			}
		})
	}
}

type states []bool

func (s states) isBadVersion(version int) bool {
	return s[version-1]
}

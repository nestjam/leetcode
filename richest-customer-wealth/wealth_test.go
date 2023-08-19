package richestcustomerwealth

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		accounts [][]int
		want     int
	}{
		{
			desc:     "single customer has one account",
			accounts: [][]int{{0}},
			want:     0,
		},
		{
			desc:     "single customer has two accounts",
			accounts: [][]int{{1, 1}},
			want:     2,
		},
		{
			desc:     "two customers have one account",
			accounts: [][]int{{2}, {1}},
			want:     2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := maximumWealth(tC.accounts)

			if got != tC.want {
				t.Errorf("got %v but want %v", got, tC.want)
			}
		})
	}
}

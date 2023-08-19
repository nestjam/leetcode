package numberstepstoreducenumbertozero

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		num int
		want int
	}{
		{
			desc: "0",
			num: 0,
			want: 0,
		},
		{
			desc: "1",
			num: 1,
			want: 1,
		},
		{
			desc: "2",
			num: 2,
			want: 2,
		},
		{
			desc: "3",
			num: 3,
			want: 3,
		},
		{
			desc: "14",
			num: 14,
			want: 6,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := numberOfSteps(tC.num)

			if got != tC.want {
				t.Errorf("expected %v got %v", tC.want, got)
			}
		})
	}
}

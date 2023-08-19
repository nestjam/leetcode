package romantointeger

import "testing"

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		roman string
		want int
	}{
		{
			roman: "I",
			want: 1,
		},
		{
			roman: "II",
			want: 2,
		},
		{
			roman: "III",
			want: 3,
		},
		{
			roman: "IV",
			want: 4,
		},
		{
			roman: "V",
			want: 5,
		},
		{
			roman: "IX",
			want: 9,
		},
		{
			roman: "X",
			want: 10,
		},
		{
			roman: "XL",
			want: 40,
		},
		{
			roman: "L",
			want: 50,
		},
		{
			roman: "XC",
			want: 90,
		},
		{
			roman: "C",
			want: 100,
		},
		{
			roman: "CD",
			want: 400,
		},
		{
			roman: "D",
			want: 500,
		},
		{
			roman: "CM",
			want: 900,
		},
		{
			roman: "M",
			want: 1000,
		},
		{
			roman: "MCMXCIV",
			want: 1994,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.roman, func(t *testing.T) {
			got := romanToInt(tC.roman)

			if got != tC.want {
				t.Errorf("expected %v but %v", tC.want, got)
			}
		})
	}
}

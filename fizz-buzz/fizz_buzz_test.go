package fizzbuzz

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		n    int
		want []string
	}{
		{
			desc: "1",
			n:    1,
			want: []string{"1"},
		},
		{
			desc: "2",
			n:    2,
			want: []string{"1", "2"},
		},
		{
			desc: "2",
			n:    2,
			want: []string{"1", "2"},
		},
		{
			desc: "Fizz",
			n:    3,
			want: []string{"1", "2", "Fizz"},
		},
		{
			desc: "Buzz",
			n:    5,
			want: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			desc: "FizzBuzz",
			n:    15,
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := fizzBuzz(tC.n)

			if !reflect.DeepEqual(got, tC.want) {
				t.Errorf("expected %v but got %v", tC.want, got)
			}
		})
	}
}

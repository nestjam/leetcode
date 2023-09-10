package design

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	testCases := []struct {
		array []int
		want  []int
		desc  string
	}{
		{
			array: []int{1},
			want:  []int{1},
		},
		{
			array: []int{1, 2, 3},
			want:  []int{3, 2, 1},
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d %v", i+1, tC.desc), func(t *testing.T) {
			s := Constructor(tC.array)
			s.rand = rand.New(rand.NewSource(99))
			got := s.Shuffle()

			if !reflect.DeepEqual(tC.want, got) {
				t.Errorf("expected %v but %v", tC.want, got)
			}
		})
	}
}

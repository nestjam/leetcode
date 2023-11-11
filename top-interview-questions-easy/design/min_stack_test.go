package design

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPushAndPop(t *testing.T) {
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
			array: []int{1, 2},
			want:  []int{2, 1},
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d %v", i+1, tC.desc), func(t *testing.T) {
			s := NewMinStack()
			for _, v := range tC.array {
				s.Push(v)
			}

			got := make([]int, len(tC.want))
			for i := range tC.array {
				got[i] = s.Top()
				s.Pop()
			}

			if !reflect.DeepEqual(tC.want, got) {
				t.Errorf("expected %v but %v", tC.want, got)
			}
		})
	}
}

func TestGetMin(t *testing.T) {
	testCases := []struct {
		array []int
		want  []int
		desc  string
	}{
		{
			array: []int{1, 2},
			want:  []int{1, 1},
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d %v", i+1, tC.desc), func(t *testing.T) {
			s := NewMinStack()
			for i, v := range tC.array {
				s.Push(v)

				got := s.GetMin()

				if got != tC.want[i] {
					t.Errorf("expected %v but %v", tC.want, got)
				}
			}

			for i := len(tC.array) - 1; i >= 0; i-- {
				got := s.GetMin()

				if got != tC.want[i] {
					t.Errorf("expected %v but %v", tC.want, got)
				}

				s.Top()
				s.Pop()
			}
		})
	}
}

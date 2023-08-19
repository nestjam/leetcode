package kweakestrowsinmatrix

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		k int
		matrix [][]int
		want []int
	}{
		{
			desc: "1",
			k: 1,
			matrix: [][]int {
				{1, 0},
			},
			want: []int {0},
		},
		{
			desc: "2",
			k: 1,
			matrix: [][]int {
				{1, 0},
				{0, 0},
			},
			want: []int {1},
		},
		{
			desc: "2",
			k: 1,
			matrix: [][]int {
				{1, 1},
				{1, 0},
				{0, 0},
			},
			want: []int {2},
		},
		{
			desc: "3",
			k: 1,
			matrix: [][]int {
				{1, 1},
				{1, 0},
				{1, 0},
			},
			want: []int {1},
		},
		{
			desc: "4",
			k: 2,
			matrix: [][]int {
				{1, 0},
				{1, 0},
			},
			want: []int {0, 1},
		},
		{
			desc: "5",
			k: 3,
			matrix: [][]int {
				{1,1,0,0,0},
				{1,1,1,1,0},
				{1,0,0,0,0},
				{1,1,0,0,0},
				{1,1,1,1,1},
			},
			want: []int {2,0,3},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := kWeakestRows(tC.matrix, tC.k)

			if !reflect.DeepEqual(tC.want, got) {
				t.Errorf("expected '%v' but got '%v'", tC.want, got)
			}
		})
	}
}

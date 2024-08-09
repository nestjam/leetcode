package kthlargestelementinastream

import "testing"

func TestKthLargest_Add(t *testing.T) {
	type args struct {
		k    int
		val  int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				k:    1,
				nums: []int{},
				val:  2,
			},
			want: 2,
		},
		{
			args: args{
				k:    1,
				nums: []int{2},
				val:  1,
			},
			want: 2,
		},
		{
			args: args{
				k:    2,
				nums: []int{3, 2, 1},
				val:  4,
			},
			want: 3,
		},
		{
			args: args{
				k:    3,
				nums: []int{4, 5, 8, 2},
				val:  3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := Constructor(tt.args.k, tt.args.nums)
			if got := sut.Add(tt.args.val); got != tt.want {
				t.Errorf("KthLargest.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

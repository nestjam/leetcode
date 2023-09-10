package design

import (
	"math/rand"
	"time"
)

type Solution struct {
	nums []int
	shuffled []int
	rand *rand.Rand
}

func Constructor(nums []int) Solution {
	return Solution{
		nums,
		make([]int, len(nums)),
		rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (s *Solution) Reset() []int {
	return s.nums
}

func (s *Solution) Shuffle() []int {
	perm := s.rand.Perm(len(s.nums))
	for i, v := range s.nums {
		s.shuffled[perm[i]] = v
	}
	return s.shuffled
}

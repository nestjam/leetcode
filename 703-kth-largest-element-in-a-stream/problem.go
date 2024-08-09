package kthlargestelementinastream

type KthLargest struct {
	n []int
}

func Constructor(k int, nums []int) KthLargest {
	n := make([]int, 0, k)

	for i := 0; i < len(nums); i++ {
		n = add(n, nums[i])
	}

	return KthLargest{
		n,
	}
}

func (s *KthLargest) Add(val int) int {
	s.n = add(s.n, val)
	return s.n[len(s.n)-1]
}

func add(n []int, v int) []int {
	if len(n) < cap(n) {
		n = append(n, v)
		sort(n)
		return n
	}

	last := len(n) - 1
	if n[last] > v {
		return n
	}

	n[last] = v
	sort(n)

	return n
}

func sort(n []int) {
	for i := len(n) - 1; i > 0; i-- {
		if n[i] <= n[i-1] {
			break
		}

		n[i-1], n[i] = n[i], n[i-1]
	}
}

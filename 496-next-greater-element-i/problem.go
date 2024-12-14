package nextgreaterelementi

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	s := stack{}
	greater := make(map[int]int, len(nums2))

	for i := 0; i < len(nums2); i++ {
		for !s.isEmpty() {
			if s.top() >= nums2[i] {
				break
			}

			greater[s.pop()] = nums2[i]
		}
		s.push(nums2[i])
	}

	res := make([]int, len(nums1))

	for i := 0; i < len(nums1); i++ {
		res[i] = -1
		if v, ok := greater[nums1[i]]; ok {
			res[i] = v
		}
	}

	return res
}

type stack struct {
	items []int
}

func (s *stack) push(data int) {
	s.items = append(s.items, data)
}

func (s *stack) top() int {
	return s.items[len(s.items)-1]
}

func (s *stack) pop() int {
	if s.isEmpty() {
		panic("stack is empty")
	}

	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

func (s *stack) isEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

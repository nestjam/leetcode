package nrepeatedelementinsize2narray

func repeatedNTimes(nums []int) int {
	m := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return nums[i]
		}

		m[nums[i]] = struct{}{}
	}

	panic("not found")
}

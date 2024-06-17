package nrepeatedelementinsize2narray

func repeatedNTimes(nums []int) int {
	m := make(map[int]struct{}, len(nums)/4)

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return nums[i]
		}

		m[nums[i]] = struct{}{}
	}

	panic("not found")
}

package singlenumber

func singleNumber(nums []int) int {
	m := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		v := false

		if _, ok := m[nums[i]]; ok {
			v = true
		}

		m[nums[i]] = v
	}

	for k, v := range m {
		if v == false {
			return k
		}
	}

	panic("oops!")
}

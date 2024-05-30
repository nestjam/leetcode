package numberofgoodpairs

func numIdenticalPairs(nums []int) int {
	m := [101]int{}

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	n := 0

	for i := 1; i < len(m); i++ {
		if m[i] < 2 {
			continue
		}
		n += ((m[i] - 1) * m[i]) / 2
	}

	return n
}

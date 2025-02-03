package majorityelement

func majorityElement(nums []int) int {
	counts := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		counts[nums[i]]++
	}

	majority := nums[0]
	max := counts[majority]

	for k, v := range counts {
		if v > max {
			max = v
			majority = k
		}
	}

	return majority
}

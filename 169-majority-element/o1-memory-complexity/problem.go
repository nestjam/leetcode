package o1memorycomplexity

func majorityElement(nums []int) int {
	candidate := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}

		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}

	return candidate
}

package array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	j, k := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[j] == nums[i] {
			continue
		}

		j = i
		k++
		nums[k] = nums[i]
	}
	return k + 1
}

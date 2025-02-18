package sortcolors

func sortColors(nums []int) {
	var min int

	for i := 0; i < len(nums); i++ {
		min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}

		nums[i], nums[min] = nums[min], nums[i]
	}
}

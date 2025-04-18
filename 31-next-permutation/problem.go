package nextpermutation

import "slices"

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		k := -1

		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				if k == -1 {
					k = j
				} else if nums[k] > nums[j] {
					k = j
				}
			}
		}

		if k != -1 {
			nums[k], nums[i] = nums[i], nums[k]

			sortAscensing(nums[i+1:])
			return
		}
	}

	sortAscensing(nums)
}

func sortAscensing(nums []int) {
	slices.SortFunc(nums, func(a, b int) int {
		return a - b
	})
}

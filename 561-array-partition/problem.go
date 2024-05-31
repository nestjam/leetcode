package arraypartition

import "slices"

func arrayPairSum(nums []int) int {
	slices.Sort(nums)

	n := 0

	for i := 0; i < len(nums); i += 2 {
		n += nums[i]
	}

	return n
}

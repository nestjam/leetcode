package minimumsizesubarraysum

import "math"

func minSubArrayLen(target int, nums []int) int {
	sum, min := 0, math.MaxInt
	found := false
	i := 0

	for j := 0; j < len(nums); j++ {
		sum += nums[j]

		for sum >= target {
			if min > j-i+1 {
				min = j - i + 1
				found = true
			}

			sum -= nums[i]
			i++
		}
	}

	if found {
		return min
	}

	return 0
}

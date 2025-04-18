package countnumberofnicesubarrays

func numberOfSubarrays(nums []int, k int) int {
	subarraysCount := 0

	for i := 0; i < len(nums); i++ {
		count := 0

		for j := i; j < len(nums); j++ {
			if nums[j]%2 == 1 {
				count++
			}

			if count == k {
				subarraysCount++
			}

			if count > k {
				break
			}
		}
	}

	return subarraysCount
}

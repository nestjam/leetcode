package sortarraybyparity

func sortArrayByParity(nums []int) []int {
	n := 0
	
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			continue
		}
		nums[n], nums[i] = nums[i], nums[n]
		n++
	}

	return nums
}

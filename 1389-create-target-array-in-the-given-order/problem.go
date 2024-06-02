package createtargetarrayinthegivenorder

func createTargetArray(nums []int, index []int) []int {
	target := make([]int, len(nums))

	for i := 0; i < len(index); i++ {
		for j := 0; j < i; j++ {
			if index[j] < index[i] {
				continue
			}
			index[j]++
		}
	}

	for i := 0; i < len(index); i++ {
		target[index[i]] = nums[i]
	}

	return target
}

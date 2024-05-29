package howmanynumbersaresmallerthanthecurrentnumber

func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		n := 0

		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			if nums[j] < nums[i] {
				n++
			}
		}

		res[i] = n
	}

	return res
}

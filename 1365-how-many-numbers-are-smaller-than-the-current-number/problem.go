package howmanynumbersaresmallerthanthecurrentnumber

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				res[i]++
			}
			if nums[i] < nums[j] {
				res[j]++
			}
		}
	}

	return res
}

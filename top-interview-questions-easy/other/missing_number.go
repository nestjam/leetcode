package other

func missingNumber(nums []int) int {
	l := len(nums)
	sum := l * (l + 1) / 2
	for i := 0; i < l; i++ {
		sum -= nums[i]
	}
	return sum
}

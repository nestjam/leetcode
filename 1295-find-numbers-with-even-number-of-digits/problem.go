package findnumberswithevennumberofdigits

func findNumbers(nums []int) int {
	n := 0

	for i := 0; i < len(nums); i++ {
		even := false

		for nums[i] > 9 {
			nums[i] = nums[i] / 10
			even = !even
		}

		if even == true {
			n++
		}
	}

	return n
}

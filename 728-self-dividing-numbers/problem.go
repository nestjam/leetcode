package selfdividingnumbers

func selfDividingNumbers(left int, right int) []int {
	nums := make([]int, 0, 5000)

	for i := left; i <= right; i++ {
		if isSelfDividingNumber(i) {
			nums = append(nums, i)
		}
	}

	return nums
}

func isSelfDividingNumber(n int) bool {
	for i := n; i > 0; {
		m := i % 10
		if m == 0 {
			return false
		}
		if n%m > 0 {
			return false
		}

		i = i / 10
	}

	return true
}

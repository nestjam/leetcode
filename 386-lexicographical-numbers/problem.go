package lexicographicalnumbers

func lexicalOrder(n int) []int {
	nums := make([]int, n)
	i := 0

	for j := 1; j < 10; j++ {
		next(j, n, nums, &i)
	}

	return nums
}

func next(v, n int, nums []int, i *int) {
	if v > n {
		return
	}

	nums[*i] = v
	*i += 1

	v *= 10
	for j := 0; j < 10; j++ {
		next(v+j, n, nums, i)
	}
}

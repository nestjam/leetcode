package array

func rotate(nums []int, k int) {
	//rotate_solution1(nums, k)
	//rotate_solution2(nums, k) //быстрее варианта 1
	//rotate_solution3(nums, k) // комбинацияя 1 и 2
	rotate_solution4(nums, k) // reverse
}

func rotate_solution1(nums []int, k int) {
	l := len(nums)
	for i := 0; i < k; i++ {
		t := nums[l-1]
		for j := l - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = t
	}
}

func rotate_solution2(nums []int, k int) {
	l := len(nums)
	if k > l {
		k = k % l
	}

	if k == 0 {
		return
	}

	t := make([]int, k, l)
	copy(t, nums[l-k:])
	t = append(t, nums[:l-k]...)
	copy(nums, t)
}

func rotate_solution3(nums []int, k int) {
	l := len(nums)
	if k > l {
		k = k % l
	}

	if k == 0 {
		return
	}

	t := make([]int, k)
	copy(t, nums[l-k:])

	for j := l - 1; j >= k; j-- {
		nums[j] = nums[j-k]
	}

	copy(nums, t)
}

func rotate_solution4(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

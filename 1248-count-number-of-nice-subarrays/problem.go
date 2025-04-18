package countnumberofnicesubarrays

func numberOfSubarrays(nums []int, k int) int {
	count, n := 0, 0
	odds := make([]int, 0)
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			odds = append(odds, i)
		}

		if len(odds) > k {
			odds, j, n = countSubarrays(odds, i, j, k)
			count += n
		}
	}

	if len(odds) == k {
		_, _, n = countSubarrays(odds, len(nums)-1, j, k)
		count += n
	}

	return count
}

func countSubarrays(odds []int, i, j, k int) ([]int, int, int) {
	l := odds[0] - j
	r := 0
	n := 0
	if len(odds) > k {
		r = odds[len(odds)-1] - odds[len(odds)-2] - 1
	} else {
		r = i - odds[len(odds)-1]
	}

	if l > 0 {
		l++
	}
	if r > 0 {
		r++
	}

	if l > 0 && r > 0 {
		n = l * r
	} else {
		n = max(l, r, 1)
	}

	j = odds[0] + 1
	odds = odds[1:]
	return odds, j, n
}

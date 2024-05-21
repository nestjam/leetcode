package binarysearch

func search(a []int, x int) int {
	return reqursive_binary_search(a, 0, len(a)-1, x)
}

func reqursive_binary_search(a []int, p, r, x int) int {
	if p > r {
		return -1
	}

	q := (p + r) / 2

	if a[q] == x {
		return q
	} else if a[q] < x {
		return reqursive_binary_search(a, q+1, r, x)
	} else {
		return reqursive_binary_search(a, p, q-1, x)
	}
}

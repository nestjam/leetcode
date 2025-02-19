package bubblesort

func heightChecker(heights []int) int {
	arr := make([]int, len(heights))
	copy(arr, heights)

	hasSwapped := true

	for hasSwapped {
		hasSwapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				hasSwapped = true
			}
		}
	}

	n := 0

	for i := 0; i < len(heights); i++ {
		if heights[i] != arr[i] {
			n++
		}
	}

	return n
}

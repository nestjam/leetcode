package heightchecker

func heightChecker(heights []int) int {
	counts := make(map[int]byte)
	height := heights[0]

	for i := 0; i < len(heights); i++ {
		counts[heights[i]]++
		if heights[i] < height {
			height = heights[i]
		}
	}

	n := 0

	for i := 0; i < len(heights); i++ {
		for {
			if count, ok := counts[height]; ok && count > 0 {
				break
			}
			height++
		}

		if height != heights[i] {
			n++
		}
		counts[height]--
	}

	return n
}

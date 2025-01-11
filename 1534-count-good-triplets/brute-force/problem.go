package countgoodtriplets

func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if abs(arr[i]-arr[j]) > a {
				continue
			}

			for k := j + 1; k < len(arr); k++ {
				if abs(arr[j]-arr[k]) > b {
					continue
				}
				if abs(arr[i]-arr[k]) > c {
					continue
				}
				count++
			}
		}
	}

	return count
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

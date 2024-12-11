package sumofalloddlengthsubarrays

func sumOddLengthSubarrays(arr []int) int {
	sum := 0
	L := len(arr)
	h := L / 2
	isOdd := L%2 == 1

	for l := 1; l <= L; l += 2 {
		n := 1
		for i := 0; i < h; i++ {
			if i < l && i+l-1 < L {
				n = i + 1
			}
			sum += n * (arr[i] + arr[L-i-1])
		}

		if isOdd {
			if h == l-1 && h+l-1 < L {
				n = l
			}
			sum += n * arr[h]
		}
	}

	return sum
}

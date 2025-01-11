package countgoodtriplets

func countGoodTriplets(arr []int, a int, b int, c int) int {
	ij := make(map[int][]int)
	jk := make(map[int][]int)
	ik := make(map[int]map[int]struct{})

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			v := arr[i] - arr[j]
			if v < 0 {
				v = -v
			}

			if v <= a {
				ij[i] = append(ij[i], j)
			}
			if v <= b {
				jk[i] = append(jk[i], j)
			}
			if v <= c {
				if _, ok := ik[i]; !ok {
					ik[i] = make(map[int]struct{})
				}
				ik[i][j] = struct{}{}
			}
		}
	}

	count := 0

	for i := range ij {
		for _, j := range ij[i] {
			for _, k := range jk[j] {
				if v, ok := ik[i]; ok {
					if _, ok := v[k]; ok {
						count++
					}
				}
			}
		}
	}

	return count
}

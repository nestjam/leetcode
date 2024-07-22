package checkiftwostringarraysareequivalent

func arrayStringsAreEqual(w1 []string, w2 []string) bool {
	i1, j1, i2, j2 := 0, 0, 0, 0

	for i1 < len(w1) || i2 < len(w2) {
		var v1, v2 byte

		if i1 < len(w1) {
			v1 = w1[i1][j1]
			j1++

			if j1 == len(w1[i1]) {
				i1++
				j1 = 0
			}
		}

		if i2 < len(w2) {
			v2 = w2[i2][j2]
			j2++

			if j2 == len(w2[i2]) {
				i2++
				j2 = 0
			}
		}

		if v1 != v2 {
			return false
		}
	}

	return true
}

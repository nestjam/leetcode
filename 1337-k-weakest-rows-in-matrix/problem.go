package kweakestrowsinmatrix

func kWeakestRows(mat [][]int, k int) []int {
	sums := sumRows(mat)
	weakest := newWeakest(k)
	for i := 0; i < len(mat); i++ {
		weakest.add(i, sums)
	}
	return weakest
}

// func newWeakest(k int) []int {
// 	weakest := make([]int, k)
// 	for i := 0; i < k; i++ {
// 		weakest[i] = -1
// 	}
// 	return weakest
// }

// func add(ind int, weakest []int, sums []int) {
// 	for i := 0; i < len(weakest); i++ {
// 		if weakest[i] == -1 {
// 			weakest[i] = i
// 			break
// 		} else if sums[ind] < sums[weakest[i]] || (sums[ind] == sums[weakest[i]] && ind < weakest[i]) {
// 			for j := i; j < len(weakest); j++ {
// 				ind, weakest[j] = weakest[j], ind
// 			}
// 			break
// 		}
// 	}
// }

func sumRows(mat [][]int) []int {
	s := make([]int, len(mat))
	for i := 0; i < len(mat); i++ {
		s[i] = sum(mat[i])
	}
	return s
}

func sum(row []int) int {
	s := 0
	for i := 0; i < len(row); i++ {
		s += row[i]
	}
	return s
}

type weakest []int

func newWeakest(k int) weakest {
	weakest := make([]int, k)
	for i := 0; i < k; i++ {
		weakest[i] = -1
	}
	return weakest
}

func (w weakest) add(ind int, sums []int) {
	for i := 0; i < len(w); i++ {
		if w[i] == -1 {
			w[i] = i
			break
		} else if sums[ind] < sums[w[i]] || (sums[ind] == sums[w[i]] && ind < w[i]) {
			for j := i; j < len(w); j++ {
				ind, w[j] = w[j], ind
			}
			break
		}
	}
}
package array

func rotateImage(m [][]int) {
	rank := len(m)
	if rank == 1 {
		return
	}

	var r, i, j, k, ik, jk int
	for r = 0; r < rank/2; r++ {
		i = r
		j = rank - r - 1
		for k = 0; k < j-i; k++ {
			ik = i + k
			jk = j - k
			m[i][ik], m[ik][j], m[j][jk], m[jk][i] = m[jk][i], m[i][ik], m[ik][j], m[j][jk]
		}
	}
}

// func rotateImage(m [][]int) {
// 	rank := len(m)
// 	for r := 0; r < rank / 2; r++ {
// 		i := r
// 		j := rank - r - 1
// 		for k := 0; k < j - i; k++ {
// 			m[i][i + k], m[i + k][j], m[j][j - k], m[j - k][i] = m[j - k][i], m[i][i + k], m[i + k][j], m[j][j - k]
// 		}
// 	}
// }

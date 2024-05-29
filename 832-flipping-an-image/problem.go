package flippinganimage

func flipAndInvertImage(image [][]int) [][]int {
	n := len(image)
	m := len(image[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m/2; j++ {
			image[i][j], image[i][m-j-1] = image[i][m-j-1], image[i][j]
		}
		for j := 0; j < m; j++ {
			if image[i][j] == 1 {
				image[i][j] = 0
			} else {
				image[i][j] = 1
			}
		}
	}

	return image
}

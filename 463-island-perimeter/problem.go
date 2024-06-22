package islandperimeter

func islandPerimeter(grid [][]int) int {
	p := 0
	n := len(grid)
	m := len(grid[0])

	// left
	for i := 0; i < n; i++ {
		p += grid[i][0]
	}

	// middle
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			if grid[i][j-1] == 0 {
				p += grid[i][j]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m-1; j++ {
			if grid[i][j+1] == 0 {
				p += grid[i][j]
			}
		}
	}

	// right
	for i := 0; i < n; i++ {
		p += grid[i][m-1]
	}

	// top
	for j := 0; j < m; j++ {
		p += grid[0][j]
	}

	// middle
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i-1][j] == 0 {
				p += grid[i][j]
			}
		}
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < m; j++ {
			if grid[i+1][j] == 0 {
				p += grid[i][j]
			}
		}
	}

	// bottom
	for j := 0; j < m; j++ {
		p += grid[n-1][j]
	}

	return p
}

package countunguardedcellsinthegrid

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	const (
		wallAlias  = 5
		guardAlias = 6
	)

	count := m*n - len(guards) - len(walls)

	grid := make([][]int8, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int8, n)
	}

	for i := 0; i < len(walls); i++ {
		w := walls[i]
		grid[w[0]][w[1]] = wallAlias
	}

	for i := 0; i < len(guards); i++ {
		g := guards[i]
		grid[g[0]][g[1]] = guardAlias
	}

	for i := 0; i < m; i++ {
		isGuarded := false
		for j := 0; j < n; j++ {
			if grid[i][j] == guardAlias {
				isGuarded = true
			} else if grid[i][j] == wallAlias {
				isGuarded = false
			} else {
				if isGuarded {
					if grid[i][j] == 0 {
						count--
					}
					grid[i][j]++
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		isGuarded := false
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == guardAlias {
				isGuarded = true
			} else if grid[i][j] == wallAlias {
				isGuarded = false
			} else {
				if isGuarded {
					if grid[i][j] == 0 {
						count--
					}
					grid[i][j]++
				}
			}
		}
	}

	for j := 0; j < n; j++ {
		isGuarded := false
		for i := 0; i < m; i++ {
			if grid[i][j] == guardAlias {
				isGuarded = true
			} else if grid[i][j] == wallAlias {
				isGuarded = false
			} else {
				if isGuarded {
					if grid[i][j] == 0 {
						count--
					}
					grid[i][j]++
				}
			}
		}
	}

	for j := 0; j < n; j++ {
		isGuarded := false
		for i := m - 1; i >= 0; i-- {
			if grid[i][j] == guardAlias {
				isGuarded = true
			} else if grid[i][j] == wallAlias {
				isGuarded = false
			} else {
				if isGuarded {
					if grid[i][j] == 0 {
						count--
					}
					grid[i][j]++
				}
			}
		}
	}

	return count
}

package countunguardedcellsinthegrid

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	const (
		unguarded = 0
		guarded   = 1
		wall      = 5
		guard     = 6
	)

	count := m*n - len(guards) - len(walls)

	grid := make([][]int8, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int8, n)
	}

	for i := 0; i < len(walls); i++ {
		w := walls[i]
		grid[w[0]][w[1]] = wall
	}

	for i := 0; i < len(guards); i++ {
		g := guards[i]
		grid[g[0]][g[1]] = guard
	}

	for i := 0; i < m; i++ {
		isGuarded := false
		for j := 0; j < n; j++ {
			if grid[i][j] == guard {
				isGuarded = true
			} else if grid[i][j] == wall {
				isGuarded = false
			} else {
				if isGuarded {
					if grid[i][j] == unguarded {
						count--
					}
					grid[i][j] = guarded
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		isGuarded := false
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == guard {
				isGuarded = true
			} else if grid[i][j] == wall {
				isGuarded = false
			} else {
				if isGuarded {
					if grid[i][j] == unguarded {
						count--
					}
					grid[i][j] = guarded
				}
			}
		}
	}

	for j := 0; j < n; j++ {
		isGuarded := false
		for i := 0; i < m; i++ {
			if grid[i][j] == guard {
				isGuarded = true
			} else if grid[i][j] == wall {
				isGuarded = false
			} else {
				if isGuarded {
					if grid[i][j] == unguarded {
						count--
					}
					grid[i][j] = guarded
				}
			}
		}
	}

	for j := 0; j < n; j++ {
		isGuarded := false
		for i := m - 1; i >= 0; i-- {
			if grid[i][j] == guard {
				isGuarded = true
			} else if grid[i][j] == wall {
				isGuarded = false
			} else {
				if isGuarded {
					if grid[i][j] == unguarded {
						count--
					}
					grid[i][j] = guarded
				}
			}
		}
	}

	return count
}

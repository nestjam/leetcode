package dynamicprogramming

func climbStairs(n int) int {
    m := make(map[int]int, 2)
	m[1] = 1
	m[2] = 2
	return step(n, m)
}

func step(n int, m map[int]int) int {
	if _, ok := m[n]; !ok {
		m[n] = step(n-2, m) + step(n-1, m)
	}
	return m[n]
}
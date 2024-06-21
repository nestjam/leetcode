package fibonaccinumber

func fib(n int) int {
	m := make(map[int]int, n)
	m[0] = 0
	m[1] = 1
	calc(n, m)
	return m[n]
}

func calc(n int, m map[int]int) {
	if _, ok := m[n]; !ok {
		calc(n-1, m)
		calc(n-2, m)
		m[n] = m[n-1] + m[n-2]
	}
}

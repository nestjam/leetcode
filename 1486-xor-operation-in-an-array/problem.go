package xoroperationinanarray

func xorOperation(n int, start int) int {
	r := start
	for i := 1; i < n; i++ {
		r = r ^ (start + i << 1)
	}
	return r
}

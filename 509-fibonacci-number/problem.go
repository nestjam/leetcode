package fibonaccinumber

func fib(n int) int {
	switch n {
	case 1:
		return 1
	case 0:
		return 0
	default:
		return fib(n-1) + fib(n-2)
	}
}

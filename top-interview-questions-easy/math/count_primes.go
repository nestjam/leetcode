package math

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	a := make([]bool, n-2)
	for i := 2; i*i < n; i++ {
		if !a[i-2] {
			for j := i * i; j < n; j += i {
				a[j-2] = true
			}
		}
	}

	c := 0
	for i := 0; i < len(a); i++ {
		if !a[i] {
			c++
		}
	}
	return c
}

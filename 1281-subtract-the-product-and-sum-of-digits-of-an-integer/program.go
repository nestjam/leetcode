package subtracttheproductandsumofdigitsofaninteger

func subtractProductAndSum(n int) int {
	p := 1
	s := 0

	for ; n > 0; n = n / 10 {
		r := n % 10
		p *= r
		s += r
	}

	return p - s
}

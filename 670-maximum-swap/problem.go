package maximumswap

import "slices"

func maximumSwap(num int) int {
	d := parse(num)
	var swapped bool

	for i := 0; i < len(d)-1; i++ {
		j := maxIndex(d, i)

		if j != i {
			d[i], d[j] = d[j], d[i]
			swapped = true
			break
		}
	}

	if !swapped {
		return num
	}

	return join(d)
}

func maxIndex(a []int, s int) int {
	i := s
	n := a[i]

	for j := i + 1; j < len(a); j++ {
		if a[j] > n || (i != s && a[j] == n) {
			n = a[j]
			i = j
		}
	}

	return i
}

func parse(n int) []int {
	d := []int{}

	for n > 0 {
		d = append(d, n%10)
		n = n / 10
	}

	slices.Reverse(d)

	return d
}

func join(d []int) int {
	n := d[0]

	for i := 1; i < len(d); i++ {
		n = 10*n + d[i]
	}

	return n
}

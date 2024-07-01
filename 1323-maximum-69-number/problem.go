package maximum69number

func maximum69Number(num int) int {
	d := make([]bool, 0, 4)

	i := 0
	for num > 0 {
		d = append(d, num%10 == 9)
		num = num / 10
		i++
	}

	for i = len(d) - 1; i >= 0; i-- {
		if d[i] == false {
			d[i] = true
			break
		}
	}

	num = 0
	for i = len(d) - 1; i >= 0; i-- {
		num = num * 10
		if d[i] {
			num += 9
		} else {
			num += 6
		}
	}

	return num
}

package robotreturntoorigin

func judgeCircle(moves string) bool {
	var x, y int16 = 0, 0

	for i := 0; i < len(moves); i++ {
		m := moves[i]
		if m == 76 { // L
			x--
		} else if m == 82 { // R
			x++
		} else if m == 85 { // U
			y--
		} else if m == 68 { // D
			y++
		}
	}

	return x == 0 && y == 0
}

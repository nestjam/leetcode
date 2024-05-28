package splitastringinbalancedstrings

func balancedStringSplit(s string) int {
	var n int16
	var j int16

	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			j++
		} else if s[i] == 'R' {
			j--
		}

		if j == 0 {
			n++
		}
	}

	return int(n)
}

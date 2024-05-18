package defanginganipaddress

func defangIPaddr(address string) string {
	const (
		dot        = '.'
		leftBrace  = '['
		rightBrace = ']'
	)
	defanged := make([]byte, len(address)+3*2)

	j := 0
	for i := 0; i < len(address); i++ {
		c := address[i]
		if c == dot {
			defanged[j] = leftBrace
			j++
		}

		defanged[j] = c
		j++

		if c == dot {
			defanged[j] = rightBrace
			j++
		}
	}

	return string(defanged)
}
package goalparserinterpretation

func interpret(command string) string {
	b := make([]byte, 0, len(command))

	for i := 0; i < len(command); {
		if command[i] == 'G' {
			b = append(b, 'G')
			i++
		} else if command[i] == '(' && command[i+1] == ')' {
			b = append(b, 'o')
			i += 2
		} else {
			b = append(b, 'a', 'l')
			i += 4
		}
	}

	return string(b)
}

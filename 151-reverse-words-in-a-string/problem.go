package reversewordsinastring

func reverseWords(s string) string {
	var current, last *word

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if current != nil {
				current.prev, last = last, current
				current = nil
			}
			continue
		}

		if current == nil {
			current = &word{i: i, j: i}
		} else {
			current.j++
		}
	}

	if current != nil {
		current.prev, last = last, current
	}

	b := make([]byte, 0, len(s))
	i := 0
	for current = last; current != nil; current = current.prev {
		for j := current.i; j <= current.j; j++ {
			b = append(b, s[j])
			i++
		}
		if current.prev != nil {
			b = append(b, ' ')
			i++
		}
	}

	return string(b)
}

type word struct {
	i    int
	j    int
	prev *word
}

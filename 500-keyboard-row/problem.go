package keyboardrow

func findWords(words []string) []string {
	rows := make(map[byte]int8, 52)
	fillRow(rows, 0, "qwertyuiopQWERTYUIOP")
	fillRow(rows, 1, "asdfghjklASDFGHJKL")
	fillRow(rows, 2, "zxcvbnmZXCVBNM")

	oneRowWords := make([]string, 0)

	for i := 0; i < len(words); i++ {
		word := words[i]
		row := rows[word[0]]
		ok := true
		for j := 1; j < len(word); j++ {
			if rows[word[j]] != row {
				ok = false
				break
			}
		}
		if ok {
			oneRowWords = append(oneRowWords, word)
		}
	}

	return oneRowWords
}

func fillRow(rows map[byte]int8, row int8, letters string) {
	for i := 0; i < len(letters); i++ {
		rows[letters[i]] = row
	}
}

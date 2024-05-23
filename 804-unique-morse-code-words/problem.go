package uniquemorsecodewords

import "strings"

func uniqueMorseRepresentations(words []string) int {
	const maxWords = 12
	alphabet := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	set := make(map[string]struct{}, maxWords)

	for i := 0; i < len(words); i++ {
		code := encode(words[i], alphabet)
		set[code] = struct{}{}
	}

	return len(set)
}

func encode(text string, alphabet []string) string {
	const a = 97
	b := strings.Builder{}

	for i := 0; i < len(text); i++ {
		b.WriteString(alphabet[text[i]-a])
	}

	return b.String()
}

package ransomnote

const aIndexInAscii = 97

type Letters [26]int

func newLetters(text string) *Letters {
	letters := &Letters{}
	for i := 0; i < len(text); i++ {
		letters[text[i] - aIndexInAscii]++;
	}
	return letters
}

func (l *Letters) has(symbol byte) bool {
	return l[symbol - aIndexInAscii] > 0
}

func (l *Letters) useOne(symbol byte) {
	l[symbol - aIndexInAscii]--
}

func canConstruct(ransomNote string, magazine string) bool {
	letters := newLetters(magazine)

	for i := 0; i < len(ransomNote); i++ {
		if !letters.has(ransomNote[i]) {
			return false
		}
		letters.useOne(ransomNote[i])
	}
	
	return true
	
	// var letters [26]int
	
	// for i := 0; i < len(magazine); i++ {
	// 	letters[magazine[i] - aIndexInAscii]++;
	// }

	// for i := 0; i < len(ransomNote); i++ {
	// 	x := ransomNote[i] - aIndexInAscii
	// 	if letters[x] == 0 {
	// 		return false
	// 	}
	// 	letters[ransomNote[i] - aIndexInAscii]--
	// }
	
	// return true
}

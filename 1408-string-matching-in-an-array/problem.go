package stringmatchinginanarray

import (
	"slices"
	"strings"
)

func stringMatching(words []string) []string {
	slices.SortFunc(words, func(a, b string) int {
		return len(a) - len(b)
	})

	s := make([]string, 0)

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if len(words[i]) == len(words[j]) {
				continue
			}

			if strings.Contains(words[j], words[i]) {
				s = append(s, words[i])
				break
			}
		}
	}

	return s
}

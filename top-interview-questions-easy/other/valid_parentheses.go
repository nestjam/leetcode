package other

import "fmt"

const (
	openParentheses     = '('
	closeParentheses    = ')'
	openBraces          = '{'
	closeBraces         = '}'
	openSquareBrackets  = '['
	closeSquareBrackets = ']'
)

func isValid(s string) bool {
	closeBrackets := map[rune]rune{
		openParentheses:    closeParentheses,
		openBraces:         closeBraces,
		openSquareBrackets: closeSquareBrackets,
	}

	openBrackets := stack{}
	for _, r := range s {
		if isOpeningBracket(r) {
			openBrackets.push(r)
		} else if isClosingBracket(r) {
			b, _ := openBrackets.front()
			if closeBrackets[b] == r {
				openBrackets.pop()
			} else {
				return false
			}
		}
	}
	return openBrackets.empty()
}

func isOpeningBracket(b rune) bool {
	return b == openParentheses || b == openBraces || b == openSquareBrackets
}

func isClosingBracket(b rune) bool {
	return b == closeParentheses || b == closeBraces || b == closeSquareBrackets
}

type stack struct {
	stack []rune
}

func (s *stack) push(c rune) {
	s.stack = append(s.stack, c)
}

func (s *stack) pop() error {
	len := len(s.stack)
	if len > 0 {
		s.stack = s.stack[:len-1]
		return nil
	}
	return fmt.Errorf("pop error: stack is empty")
}

func (s *stack) front() (rune, error) {
	len := len(s.stack)
	if len > 0 {
		return s.stack[len-1], nil
	}
	return rune(0), fmt.Errorf("front error: stack is empty")
}

func (s *stack) empty() bool {
	return len(s.stack) == 0
}

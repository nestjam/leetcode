package replacewords

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	const space = ' '
	root := buildTrie(dictionary)

	j := 0
	newSentence := strings.Builder{}

	for i := 0; i < len(sentence); i++ {
		if sentence[i] == space {
			word := replaceWord(sentence[j:i], root)
			newSentence.WriteString(word)
			newSentence.WriteByte(space)
			j = i + 1
		}
	}

	word := replaceWord(sentence[j:], root)
	newSentence.WriteString(word)

	return newSentence.String()
}

func replaceWord(word string, node *trieNode) string {
	i := 0
	for ; i < len(word); i++ {
		if node == nil {
			continue
		}

		if child, ok := node.children[word[i]]; ok {
			node = child
			if node.isRoot {
				i++
				break
			}
		} else {
			node = nil
		}
	}

	return word[:i]
}

type trieNode struct {
	isRoot   bool
	children map[byte]*trieNode
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[byte]*trieNode),
	}
}

func buildTrie(dictionary []string) *trieNode {
	root := newTrieNode()

	for i := 0; i < len(dictionary); i++ {
		addWord(dictionary[i], root)
	}

	return root
}

func addWord(word string, node *trieNode) {
	for i := 0; i < len(word); i++ {
		if _, ok := node.children[word[i]]; !ok {
			node.children[word[i]] = newTrieNode()
		}

		node = node.children[word[i]]
	}

	node.isRoot = true
}

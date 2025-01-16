package replacewords

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	const space = ' '
	trie := buildTrie(dictionary)

	j := 0
	newSentence := strings.Builder{}

	for i := 0; i < len(sentence); i++ {
		if sentence[i] == space {
			word := replaceWord(sentence[j:i], trie)
			newSentence.WriteString(word)
			newSentence.WriteByte(space)
			j = i + 1
		}
	}

	word := replaceWord(sentence[j:], trie)
	newSentence.WriteString(word)

	return newSentence.String()
}

func replaceWord(word string, trie *trie) string {
	if prefix, ok := trie.findShortestPrefix(word); ok {
		return prefix
	}

	return word
}

func buildTrie(dictionary []string) *trie {
	trie := newTrie()

	for i := 0; i < len(dictionary); i++ {
		trie.add(dictionary[i])
	}

	return trie
}

type trie struct {
	root *trieNode
}

func newTrie() *trie {
	return &trie{
		root: newTrieNode(),
	}
}

func (t *trie) add(s string) {
	node := t.root

	for i := 0; i < len(s); i++ {
		child, ok := node.findChild(s[i])

		if !ok {
			child = newTrieNode()
			node.addChild(s[i], child)
		}

		node = child
	}

	node.isEnd = true
}

func (t *trie) findShortestPrefix(s string) (string, bool) {
	i := 0
	node := t.root

	for ; i < len(s); i++ {
		if child, ok := node.findChild(s[i]); ok {
			node = child
			if node.isEnd {
				i++
				return s[:i], true
			}
		} else {
			break
		}
	}

	return "", false
}

type trieNode struct {
	isEnd    bool
	children map[byte]*trieNode
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[byte]*trieNode),
	}
}

func (t *trieNode) findChild(c byte) (*trieNode, bool) {
	child, ok := t.children[c]
	return child, ok
}

func (t *trieNode) addChild(c byte, child *trieNode) {
	t.children[c] = child
}

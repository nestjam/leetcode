package trie

func stringMatching(words []string) []string {
	rootNode := buildTrie(words)

	substrings := make([]string, 0)

	for i := 0; i < len(words); i++ {
		if isSubstring(words[i], rootNode) {
			substrings = append(substrings, words[i])
		}
	}

	return substrings
}

func isSubstring(word string, rootNode *trieNode) bool {
	node := rootNode

	for i := 0; i < len(word); i++ {
		node = node.children[word[i]]
	}

	if node.frequency > 1 {
		return true
	}

	return false
}

type trieNode struct {
	frequency int
	children  map[byte]*trieNode
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[byte]*trieNode),
	}
}

func buildTrie(words []string) *trieNode {
	root := newTrieNode()

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			addWord(words[i][j:], root)
		}
	}

	return root
}

func addWord(word string, node *trieNode) {
	for j := 0; j < len(word); j++ {
		if _, ok := node.children[word[j]]; !ok {
			node.children[word[j]] = newTrieNode()
		}
		node = node.children[word[j]]
		node.frequency++
	}
}

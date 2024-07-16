package implementtrie

type Trie struct {
	root *trieNode
}

type trieNode struct {
	children map[byte]*trieNode
	isWord   bool
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[byte]*trieNode),
	}
}

func Constructor() Trie {
	return Trie{
		root: newTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	node := t.root

	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, ok := node.children[char]; !ok {
			child := newTrieNode()
			node.children[char] = child
		}

		node = node.children[char]
	}

	node.isWord = true
}

func (t *Trie) Search(word string) bool {
	node := t.root

	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, ok := node.children[char]; !ok {
			return false
		}

		node = node.children[char]
	}

	return node.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root

	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		if _, ok := node.children[char]; !ok {
			return false
		}

		node = node.children[char]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

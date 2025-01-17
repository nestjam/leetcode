package designaddandsearchwordsdatastructure

type WordDictionary struct {
	root *trieNode
}

func Constructor() WordDictionary {
	root := newTrieNode()
	return WordDictionary{
		root: &root,
	}
}

func (d *WordDictionary) AddWord(word string) {
	node := d.root

	for i := 0; i < len(word); i++ {
		child, ok := node.findChild(word[i])

		if !ok {
			newNode := newTrieNode()
			child = &newNode
			node.addChild(word[i], child)
		}

		node = child
	}

	node.isEnd = true
}

func (d *WordDictionary) Search(word string) bool {
	const wildChar = '.'

	nodes := []*trieNode{d.root}

	for i := 0; i < len(word); i++ {
		children := make([]*trieNode, 0)

		for j := 0; j < len(nodes); j++ {
			if word[i] == wildChar {
				children = append(children, nodes[j].getChildren()...)
				continue
			}

			if child, ok := nodes[j].findChild(word[i]); ok {
				children = append(children, child)
			}
		}

		if len(children) == 0 {
			return false
		}

		nodes = children
	}

	for i := 0; i < len(nodes); i++ {
		if nodes[i].isEnd {
			return true
		}
	}

	return false
}

type trieNode struct {
	isEnd    bool
	children map[byte]*trieNode
}

func newTrieNode() trieNode {
	return trieNode{
		children: make(map[byte]*trieNode),
	}
}

func (t *trieNode) findChild(c byte) (*trieNode, bool) {
	child, ok := t.children[c]
	return child, ok
}

func (t *trieNode) getChildren() []*trieNode {
	children := make([]*trieNode, 0, len(t.children))

	for _, child := range t.children {
		children = append(children, child)
	}

	return children
}

func (t *trieNode) addChild(c byte, child *trieNode) {
	t.children[c] = child
}

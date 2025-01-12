package mapsumpairs

type MapSum struct {
	root *trieNode
}

func Constructor() MapSum {
	root := newTrieNode()
	return MapSum{
		root: &root,
	}
}

func (m *MapSum) Insert(key string, val int) {
	node := m.findNode(key)
	if node != nil && node.val > 0 {
		val = val - node.val
	}

	node = m.root
	for i := 0; i < len(key); i++ {
		if _, ok := node.children[key[i]]; !ok {
			child := newTrieNode()
			node.children[key[i]] = &child
		}
		node = node.children[key[i]]
		node.sum += val
	}
	node.val += val
}

func (m *MapSum) Sum(prefix string) int {
	node := m.root

	for i := 0; i < len(prefix); i++ {
		if _, ok := node.children[prefix[i]]; !ok {
			return 0
		}
		node = node.children[prefix[i]]
	}

	return node.sum
}

func (m *MapSum) findNode(prefix string) *trieNode {
	node := m.root

	for i := 0; i < len(prefix); i++ {
		if _, ok := node.children[prefix[i]]; !ok {
			return nil
		}
		node = node.children[prefix[i]]
	}

	return node
}

type trieNode struct {
	sum      int
	val      int
	children map[byte]*trieNode
}

func newTrieNode() trieNode {
	return trieNode{
		children: make(map[byte]*trieNode),
	}
}

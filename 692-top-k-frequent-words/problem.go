package topkfrequentwords

import (
	"container/heap"
	"strings"
)

func topKFrequent(words []string, k int) []string {
	usings := newWordsUsing(words)

	topK := make([]string, 0, k)
	for i := 0; i < k; i++ {
		topK = append(topK, usings.next().word)
	}

	return topK
}

type wordUsing struct {
	word  string
	count int
}

type wordsUsing []*wordUsing

func newWordsUsing(words []string) *wordsUsing {
	usings := make(map[string]int)
	for i := 0; i < len(words); i++ {
		usings[words[i]]++
	}

	h := make(wordsUsing, 0)
	for k, v := range usings {
		h = append(h, &wordUsing{k, v})
	}

	heap.Init(&h)
	return &h
}

func (h wordsUsing) Len() int { return len(h) }

func (h wordsUsing) Less(i, j int) bool { 
	x := h[i]
	y := h[j]
	
	if x.count != y.count {
		return x.count > y.count 
	}
	
	return strings.Compare(x.word, y.word) < 0
}

func (h wordsUsing) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *wordsUsing) Push(x any) {
	*h = append(*h, x.(*wordUsing))
}

func (h *wordsUsing) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *wordsUsing) next() *wordUsing {
	return heap.Pop(h).(*wordUsing)
}

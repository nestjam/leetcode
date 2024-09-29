package networkdelaytime

import (
	"container/heap"
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {
	d := newDistances(n)
	d[k-1] = 0

	g := newGraph(n, times)

	q := newPathHeap()
	q.add(path{0, k - 1})

	for len(q) > 0 {
		p := q.get()

		if p.d > d[p.n] {
			continue
		}

		m := g[p.n]
		for u, w := range m {
			if p.d+w < d[u] {
				d[u] = p.d + w
				q.add(path{d[u], u})
			}
		}
	}

	max := max(d)
	if max == math.MaxInt {
		return -1
	}
	return max
}

func max(n []int) int {
	m := 0

	for i := 0; i < len(n); i++ {
		if m < n[i] {
			m = n[i]
		}
	}
	return m
}

func newDistances(n int) []int {
	d := make([]int, n)

	for i := 0; i < n; i++ {
		d[i] = math.MaxInt
	}
	return d
}

func newGraph(n int, edges [][]int) []map[int]int {
	g := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		g[i] = make(map[int]int)
	}

	for i := 0; i < len(edges); i++ {
		edge := edges[i]
		g[edge[0]-1][edge[1]-1] = edge[2]
	}
	return g
}

type path struct {
	d int
	n int
}

type pathHeap []path

func newPathHeap() pathHeap {
	h := make(pathHeap, 0)
	heap.Init(&h)
	return h
}

func (h pathHeap) Len() int { return len(h) }

func (h pathHeap) Less(i, j int) bool {
	return h[i].d < h[j].d
}

func (h pathHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *pathHeap) Push(x any) {
	*h = append(*h, x.(path))
}

func (h *pathHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *pathHeap) add(p path) {
	heap.Push(h, p)
}

func (h *pathHeap) get() path {
	return heap.Pop(h).(path)
}

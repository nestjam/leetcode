package pathwithmaximumprobability

import (
	"container/heap"
	"math"
)

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	dist := make([]float64, n)

	graph := initGraph(n, edges)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}
	dist[start_node] = 1

	q := newDistHeap(n)
	q.add(start_node, dist[start_node])

	for len(q) > 0 {
		d := q.get()

		if d.u == -1 || d.u == end_node {
			break
		}

		neighbors := graph[d.u]
		for i := 0; i < len(neighbors); i++ {
			v := neighbors[i]
			alt := dist[d.u] * succProb[v.edge]
			if alt > dist[v.node] {
				dist[v.node] = alt
				q.add(int(v.node), alt)
			}
		}
	}

	if dist[end_node] == -1 {
		return 0
	}

	return round(dist[end_node])
}

func initGraph(n int, edges [][]int) map[int]neighbors {
	g := make(map[int]neighbors, n)

	for i := 0; i < n; i++ {
		g[i] = make(neighbors, 0)
	}

	for i := 0; i < len(edges); i++ {
		start := edges[i][0]
		end := edges[i][1]
		g[start] = append(g[start], neighbor{node: int16(end), edge: int16(i)})
		g[end] = append(g[end], neighbor{node: int16(start), edge: int16(i)})
	}

	return g
}

func round(v float64) float64 {
	d := float64(100_000)
	return math.Round(v*d) / d
}

type neighbor struct {
	edge int16
	node int16
}

type neighbors []neighbor

type dist struct {
	u int
	p float64
}

type distHeap []dist

func newDistHeap(capacity int) distHeap {
	h := make(distHeap, 0, capacity)
	heap.Init(&h)
	return h
}

func (h distHeap) Len() int { return len(h) }

func (h distHeap) Less(i, j int) bool {
	return h[i].p > h[j].p
}

func (h distHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *distHeap) Push(x any) {
	*h = append(*h, x.(dist))
}

func (h *distHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *distHeap) add(u int, p float64) {
	heap.Push(h, dist{u, p})
}

func (h *distHeap) get() dist {
	return heap.Pop(h).(dist)
}

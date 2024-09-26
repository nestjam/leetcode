package kclosestpointstoorigin

import "container/heap"

func kClosest(points [][]int, k int) [][]int {
	if len(points) == k {
		return points
	}

	h := newPointHeap(points)
	kPoints := make([][]int, 0, k)

	for i := 0; i < k; i++ {
		kPoints = append(kPoints, h.next().p)
	}

	return kPoints
}

type point struct {
	p []int
	d int
}

func newPoint(p []int) point {
	d := p[0]*p[0] + p[1]*p[1]
	return point{p, d}
}

type pointHeap []point

func newPointHeap(points [][]int) pointHeap {
	h := make(pointHeap, 0, len(points))

	for _, p := range points {
		h = append(h, newPoint(p))
	}

	heap.Init(&h)
	return h
}

func (h pointHeap) Len() int { return len(h) }

func (h pointHeap) Less(i, j int) bool {
	return h[i].d < h[j].d
}

func (h pointHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *pointHeap) Push(x any) {
	*h = append(*h, x.(point))
}

func (h *pointHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *pointHeap) next() point {
	return heap.Pop(h).(point)
}

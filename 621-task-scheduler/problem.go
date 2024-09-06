package taskscheduler

import "container/heap"

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	const idle byte = 0
	flow := []byte{}

	m := make(map[byte]int)
	for i := 0; i < len(tasks); i++ {
		m[tasks[i]]++
	}

	for {
		h := newItemHeap(m)
		l := len(m)

		for i := 0; i < min(n+1, l); i++ {
			item := h.next()
			flow = append(flow, item.task)
			m[item.task]--
		}

		for k, v := range m {
			if v == 0 {
				delete(m, k)
			}
		}

		if len(m) == 0 {
			break
		}

		for i := 0; i < max(0, n-l+1); i++ {
			flow = append(flow, idle)
		}
	}

	return len(flow)
}

type item struct {
	task  byte
	count int
}

type itemHeap []*item

func newItemHeap(m map[byte]int) *itemHeap {
	h := make(itemHeap, 0, len(m))

	for k, v := range m {
		h = append(h, &item{k, v})
	}

	heap.Init(&h)
	return &h
}

func (h itemHeap) Len() int { return len(h) }

func (h itemHeap) Less(i, j int) bool { return h[i].count > h[j].count }

func (h itemHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *itemHeap) Push(x any) {
	*h = append(*h, x.(*item))
}

func (h *itemHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *itemHeap) next() *item {
	return heap.Pop(h).(*item)
}

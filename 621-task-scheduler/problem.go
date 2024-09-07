package taskscheduler

import "container/heap"

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	const idle byte = 0
	flow := make([]byte, 0, len(tasks))
	groups := groupByTask(tasks)

	for {
		orderedGroups := sortByCount(groups)

		l := len(groups)

		for i := 0; i < min(n+1, l); i++ {
			item := orderedGroups[i]
			flow = append(flow, item.task)
			item.count--
		}

		groups = deleteEmpty(groups)

		if len(groups) == 0 {
			break
		}

		for i := 0; i < max(0, n-l+1); i++ {
			flow = append(flow, idle)
		}
	}

	return len(flow)
}

func deleteEmpty(groups []*taskGroup) []*taskGroup {
	oldItems := groups

	groups = make([]*taskGroup, 0)

	for i := 0; i < len(oldItems); i++ {
		if oldItems[i].count > 0 {
			groups = append(groups, oldItems[i])
		}
	}

	return groups
}

func sortByCount(groups []*taskGroup) []*taskGroup {
	heap := newItemHeap(groups)
	orderedGroups := make([]*taskGroup, 0, len(groups))

	for heap.Len() > 0 {
		orderedGroups = append(orderedGroups, heap.next())
	}

	return orderedGroups
}

func groupByTask(tasks []byte) []*taskGroup {
	m := make(map[byte]int)

	for i := 0; i < len(tasks); i++ {
		m[tasks[i]]++
	}

	groups := make([]*taskGroup, 0, len(m))

	for k, v := range m {
		groups = append(groups, &taskGroup{k, v})
	}

	return groups
}

type taskGroup struct {
	task  byte
	count int
}

type itemHeap []*taskGroup

func newItemHeap(items []*taskGroup) *itemHeap {
	h := make(itemHeap, 0, len(items))

	for i := 0; i < len(items); i++ {
		h = append(h, items[i])
	}

	heap.Init(&h)
	return &h
}

func (h itemHeap) Len() int { return len(h) }

func (h itemHeap) Less(i, j int) bool { return h[i].count > h[j].count }

func (h itemHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *itemHeap) Push(x any) {
	*h = append(*h, x.(*taskGroup))
}

func (h *itemHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *itemHeap) next() *taskGroup {
	return heap.Pop(h).(*taskGroup)
}

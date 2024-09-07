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
		orderedGroups := sortByTaskCount(groups)
		oldCount := len(orderedGroups)
		count := oldCount

		for i := 0; i < min(n+1, oldCount); i++ {
			item := orderedGroups[i]
			flow = append(flow, item.task)
			item.count--
			if item.count == 0 {
				count--
			}
		}

		if count == 0 {
			break
		}

		for i := 0; i < max(0, n-oldCount+1); i++ {
			flow = append(flow, idle)
		}
	}

	return len(flow)
}

func sortByTaskCount(groups []*taskGroup) []*taskGroup {
	heap := newTaskGroupHeap(groups)
	orderedGroups := make([]*taskGroup, 0, heap.Len())

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

type taskGroupHeap []*taskGroup

func newTaskGroupHeap(items []*taskGroup) *taskGroupHeap {
	h := make(taskGroupHeap, 0, len(items))

	for i := 0; i < len(items); i++ {
		if items[i].count > 0 {
			h = append(h, items[i])
		}
	}

	heap.Init(&h)
	return &h
}

func (h taskGroupHeap) Len() int { return len(h) }

func (h taskGroupHeap) Less(i, j int) bool { return h[i].count > h[j].count }

func (h taskGroupHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *taskGroupHeap) Push(x any) {
	*h = append(*h, x.(*taskGroup))
}

func (h *taskGroupHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *taskGroupHeap) next() *taskGroup {
	return heap.Pop(h).(*taskGroup)
}

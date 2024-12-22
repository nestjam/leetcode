package twobestnonoverlappingevents

import (
	"slices"
)

func maxTwoEvents(intervals [][]int) int {
	events := make([]event, 0, 2*len(intervals))
	for i := 0; i < len(intervals); i++ {
		interval := intervals[i]
		events = append(events, event{time: interval[0], start: true, value: interval[2]})
		events = append(events, event{time: interval[1] + 1, start: false, value: interval[2]})
	}

	slices.SortFunc(events, func(a, b event) int {
		t := a.time - b.time

		if t != 0 {
			return t
		}

		if !a.start && b.start {
			return -1
		}
		if a.start && !b.start {
			return 1
		}
		return 0
	})

	var n, m int

	for i := 0; i < len(events); i++ {
		if events[i].start {
			n = max(n, events[i].value+m)
		} else {
			m = max(m, events[i].value)
		}
	}

	return n
}

type event struct {
	time  int
	value int
	start bool
}

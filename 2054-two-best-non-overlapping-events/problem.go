package twobestnonoverlappingevents

import (
	"slices"
)

func maxTwoEvents(events [][]int) int {
	sortByStartDate(events)

	eventsByEndTime := make([][]int, len(events))
	copy(eventsByEndTime, events)
	sortByEndDate(eventsByEndTime)

	maxSum := 0
	firstEvent := 0
	secondEvent := 0

	for i := 0; i < len(eventsByEndTime); i++ {
		if eventsByEndTime[firstEvent][1] != eventsByEndTime[i][1] {
			sum := eventsByEndTime[firstEvent][2]
			index := len(events)

			for j := secondEvent; j < len(events); j++ {
				if events[j][0] > eventsByEndTime[firstEvent][1] {
					index = j
					break
				}
			}

			secondEvent = index

			if index < len(events) {
				secondValue := events[index][2]
				for j := index+1; j < len(events); j++ {
					if secondValue < events[j][2] {
						secondValue = events[j][2]
					}
				}

				sum += secondValue
			}

			if maxSum < sum {
				maxSum = sum
			}

			firstEvent = i
		}

		if eventsByEndTime[firstEvent][2] < eventsByEndTime[i][2] {
			firstEvent = i
		}
	}

	if maxSum < eventsByEndTime[firstEvent][2] {
		maxSum = eventsByEndTime[firstEvent][2]
	}

	return maxSum
}

func sortByStartDate(events [][]int) {
	slices.SortFunc(events, func(a, b []int) int {
		return a[0] - b[0]
	})
}

func sortByEndDate(events [][]int) {
	slices.SortFunc(events, func(a, b []int) int {
		return a[1] - b[1]
	})
}

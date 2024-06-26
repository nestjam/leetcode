package bookingmeetingroom

import "math"

// https://new.contest.yandex.ru/48557/problem?id=215/2023_04_06/RdGbbmsLQn
func getMaxBookings(intervals [][]int) int {
	n := 0
	l := 0

	for {
		interval := nextInterval(l, intervals)

		if interval == nil {
			break
		}

		n++
		l = interval[1]
	}

	return n
}

func nextInterval(l int, intervals [][]int) []int {
	var interval []int
	var r int = math.MaxInt

	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > l && intervals[i][1] < r {
			r = intervals[i][1]
			interval = intervals[i]
		}
	}

	return interval
}

package mergeintervals

import "slices"

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	merged := make([][]int, 0, len(intervals))
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		latestMerged := merged[len(merged)-1]
		interval := intervals[i]

		if latestMerged[1] < interval[0] {
			merged = append(merged, intervals[i])
			continue
		}

		latestMerged[1] = max(interval[1], latestMerged[1])
	}

	return merged
}
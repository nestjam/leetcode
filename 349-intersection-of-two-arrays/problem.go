package intersectionoftwoarrays

import (
	"slices"
)

func intersection(nums1 []int, nums2 []int) []int {
	slices.Sort(nums1)
	slices.Sort(nums2)

	i, j, k := 0, 0, -1
	s := make([]int, 0, 100)

	for i < len(nums1) && j < len(nums2) {
		v := 0
		if nums1[i] == nums2[j] {
			v = nums1[i]
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
			continue
		} else {
			j++
			continue
		}

		if v > k {
			k = v
			s = append(s, v)
		}
	}

	return s
}

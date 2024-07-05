package intersectionoftwoarrays

import (
	"slices"
)

func intersection(nums1 []int, nums2 []int) []int {
	slices.Sort(nums1)
	slices.Sort(nums2)

	i, j := 0, 0
	s := make([]int, 0, 500)

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

		if len(s) == 0 || v > s[len(s)-1] {
			s = append(s, v)
		}
	}

	return s
}

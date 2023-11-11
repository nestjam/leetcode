package array

func intersect(nums1 []int, nums2 []int) []int {
	k := 0
	for i := 0; i < len(nums1); i++ {
		for j := k; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				nums2[k], nums2[j] = nums2[j], nums2[k]
				k++
				break
			}
		}
	}
	return nums2[:k]
}

/*
func intersect(nums1 []int, nums2 []int) []int {
    if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	counters := make([]uint16, 1001)
	for i := 0; i < len(nums2); i++ {
		counters[nums2[i]]++
	}
	result := make([]int, 0, len(nums1))
	for i := 0; i < len(nums1); i++ {
		if counters[nums1[i]] > 0 {
			result = append(result, nums1[i])
			counters[nums1[i]]--
		}
	}
	return result
}
*/

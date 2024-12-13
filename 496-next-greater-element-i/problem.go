package nextgreaterelementi

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	indexes := make(map[int]int, len(nums2))

	for i := 0; i < len(nums2); i++ {
		indexes[nums2[i]] = i
	}

	res := make([]int, len(nums1))

	for i := 0; i < len(nums1); i++ {
		res[i] = -1
		for j := indexes[nums1[i]] + 1; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				res[i] = nums2[j]
				break
			}
		}
	}

	return res
}

package sortingandsearching

func merge(nums1 []int, m int, nums2 []int, n int) {
	i1, i2, j := m-1, n-1, m+n-1
	for j >= 0 {
		if i1 >= 0 && i2 < 0 {
			i1--
		} else if i1 >= 0 && nums1[i1] > nums2[i2] {
			nums1[j] = nums1[i1]
			i1--
		} else {
			nums1[j] = nums2[i2]
			i2--
		}
		j--

		/*
			if i1 < 0 {
				nums1[j] = nums2[i2]
				i2--
			} else if i2 < 0 {
				i1--
			} else if nums1[i1] > nums2[i2] {
				nums1[j] = nums1[i1]
				i1--
			} else {
				nums1[j] = nums2[i2]
				i2--
			}
			j--
		*/
	}
}

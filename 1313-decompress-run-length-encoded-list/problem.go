package decompressrunlengthencodedlist

func decompressRLElist(nums []int) []int {
	n := 0
	for i := 0; i < len(nums); i += 2 {
		n += nums[i]
	}
	decoded := make([]int, n)

	l := 0
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			decoded[l] = nums[i+1]
			l++
		}
	}

	return decoded
}

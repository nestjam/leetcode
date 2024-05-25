package shufflethearray

func shuffle(nums []int, n int) []int {
	f := nums[:n]
	s := nums[n:]
	r := make([]int, len(nums))
	for i := 0; i < n; i++ {
		r[2*i] = f[i]
		r[2*i+1] = s[i]
	}
	return r
}

package maximumaveragesubarray

func findMaxAverage(nums []int, k int) float64 {
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum

	for i := 1; i < len(nums)-k+1; i++ {
		sum += nums[i+k-1] - nums[i-1]

		if maxSum < sum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

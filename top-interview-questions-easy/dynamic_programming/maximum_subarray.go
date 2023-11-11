package dynamicprogramming

func maxSubArray(nums []int) int {
	curSum := nums[0]
	maxSum := curSum
	for i := 1; i < len(nums); i++ {
		if curSum < 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}
		if maxSum < curSum {
			maxSum = curSum
		}
	}
	return maxSum
}

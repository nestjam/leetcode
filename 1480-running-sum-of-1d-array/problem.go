package runningsumof1darray

// func runningSum(nums []int) []int {
//     sum := make([]int, len(nums))
//     sum[0] = nums[0]

//     for i := 1; i < len(nums); i++ {
//         sum[i] = sum[i - 1] + nums[i]
//     }

//     return sum
// }

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	return nums
}

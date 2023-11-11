package dynamicprogramming

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	p := make([]int, len(nums))
	p[0], p[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		p[i] = max(nums[i]+p[i-2], p[i-1])
	}

	return p[len(p)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// slow...

// type interval struct {
// 	start int
// 	end   int
// }

// func (i *interval) len() int {
// 	return i.end - i.start + 1
// }

// func rob(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	m := map[interval]int{}
// 	return robHouses(nums, interval{0, len(nums) - 1}, m)
// }

// func robHouses(nums []int, s interval, m map[interval]int) int {
// 	if s.start == s.end {
// 		return nums[s.start]
// 	}

// 	if s.start == s.end - 1 {
// 		if nums[s.start] > nums[s.end] {
// 			return nums[s.start]
// 		} else {
// 			return nums[s.end]
// 		}
// 	}

// 	maxProfit := 0

// 	for i := s.start; i <= s.end; i++ {
// 		profit := nums[i]

// 		if i >= s.start + 2 {
// 			left := interval{s.start, i - 2}
// 			if left.len() > 2 {
// 				if _, ok := m[left]; !ok {
// 					m[left] = robHouses(nums, left, m)
// 				}
// 				profit += m[left]
// 			} else {
// 				profit += robHouses(nums, left, m)
// 			}
// 		}

// 		if i <= s.end - 2 {
// 			right := interval{i + 2, s.end}
// 			if right.len() > 2 {
// 				if _, ok := m[right]; !ok {
// 					m[right] = robHouses(nums, right, m)
// 				}
// 				profit += m[right]
// 			} else {
// 				profit += robHouses(nums, right, m)
// 			}
// 		}

// 		if profit > maxProfit {
// 			maxProfit = profit
// 		}
// 	}

// 	return maxProfit
// }

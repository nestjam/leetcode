package dynamicprogramming

func maxProfit(p []int) int {
	b := 0
	maxProfit := 0
	for i := 1; i < len(p); i++ {
		if p[i] < p[b] {
			b = i
		} 

		profit := p[i] - p[b]
		if maxProfit < profit {
			maxProfit = profit
		}
	}
	return maxProfit
}

// func maxProfit(p []int) int {
//     b := []int{ 0 }
// 	maxProfit := 0

// 	for i := 1; i < len(p); i++ {
// 		if p[i] < p[b[i - 1]] {
// 			b = append(b, i)
// 		} else {
// 			b = append(b, b[i - 1])
// 		}

// 		profit := p[i] - p[b[i]]
// 		if maxProfit < profit {
// 			maxProfit = profit
// 		}
// 	}
	
// 	return maxProfit
// }
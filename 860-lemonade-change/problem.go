package lemonadechange

func lemonadeChange(bills []int) bool {
	cash := make([]int, 3)

	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			cash[0]++
		} else if bills[i] == 10 {
			cash[1]++

			if cash[0] > 0 {
				cash[0]--
			} else {
				return false
			}
		} else {
			cash[2]++

			if cash[0] > 0 && cash[1] > 0 {
				cash[1]--
				cash[0]--
			} else if cash[0] > 2 && cash[1] == 0 {
				cash[0] -= 3
			} else {
				return false
			}
		}
	}

	return true
}

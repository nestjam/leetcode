package richestcustomerwealth

// func maximumWealth(accounts [][]int) int {
// 	maxWealth := 0
// 	channel := make(chan int)

// 	for _, customerAccounts := range accounts {
// 		go func(a []int) {
// 			channel <- getCustomerWealth(a)
// 		}(customerAccounts)
// 	}

// 	for i := 0; i < len(accounts); i++ {
// 		r := <-channel
// 		if r > maxWealth {
// 			maxWealth = r
// 		}
// 	}

// 	return maxWealth
// }

func maximumWealth(accounts [][]int) int {
	maxWealth := 0

	for i := 0; i < len(accounts); i++ {
		wealth := getCustomerWealth(accounts[i])
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}

	return maxWealth
}

func getCustomerWealth(accounts []int) int {
	wealth := 0

	for i := 0; i < len(accounts); i++ {
		wealth += accounts[i]
	}

	return wealth
}

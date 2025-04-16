package averagewaitingtime

func averageWaitingTime(customers [][]int) float64 {
	waitingTime := 0
	var readyTime int
	
	for i := 0; i < len(customers); i++ {
		if readyTime < customers[i][0] {
			readyTime = customers[i][0]
		}
		readyTime += customers[i][1]
		waitingTime += readyTime - customers[i][0]
	}

	return float64(waitingTime) / float64(len(customers))
}
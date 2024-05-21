package numberstepstoreducenumbertozero

func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}

	output := 1

	for num > 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num--
		}
		output++
	}
	return output
}

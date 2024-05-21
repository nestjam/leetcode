package kidswiththegreatestnumberofcandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	max := candies[0]
	for i := 1; i < n; i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}

	res := make([]bool, n)
	for i := 0; i < n; i++ {
		res[i] = (candies[i] + extraCandies) >= max
	}
	return res
}

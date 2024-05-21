package jewelsandstones

func numJewelsInStones(jewels string, stones string) int {
	set := make(map[byte]struct{}, len(jewels))
	for i := 0; i < len(jewels); i++ {
		set[jewels[i]] = struct{}{}
	}

	num := 0
	for i := 0; i < len(stones); i++ {
		if _, ok := set[stones[i]]; ok {
			num++
		}
	}

	return num
}

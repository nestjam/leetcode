package destinationcity

func destCity(paths [][]string) string {
	startCities := make(map[string]struct{}, 0)

	for i := 0; i < len(paths); i++ {
		startCities[paths[i][0]] = struct{}{}
	}

	for i := 0; i < len(paths); i++ {
		if _, ok := startCities[paths[i][1]]; !ok {
			return paths[i][1]
		}
	}

	return ""
}

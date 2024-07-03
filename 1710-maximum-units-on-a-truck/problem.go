package maximumunitsonatruck

import "slices"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	slices.SortFunc(boxTypes, func(a, b []int) int {
		return b[1] - a[1]
	})

	i := 0
	size := 0
	units := 0

	for size < truckSize && len(boxTypes) > i {
		n := min(truckSize-size, boxTypes[i][0])
		units += n * boxTypes[i][1]
		size += n
		boxTypes[i][0] -= n

		if boxTypes[i][0] == 0 {
			i++
		}
	}

	return units
}

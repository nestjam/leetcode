package maximumunitsonatruck

import "slices"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	slices.SortFunc(boxTypes, func(a, b []int) int {
		return b[1] - a[1]
	})

	size := 0
	units := 0

	for size < truckSize && len(boxTypes) > 0 {
		n := min(truckSize-size, boxTypes[0][0])
		units += n * boxTypes[0][1]
		size += n
		boxTypes[0][0] -= n

		if boxTypes[0][0] == 0 {
			boxTypes = boxTypes[1:]
		}
	}

	return units
}

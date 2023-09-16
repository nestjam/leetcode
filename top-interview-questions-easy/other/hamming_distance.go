package other

func hammingDistance(x int, y int) int {
    t := uint(x) ^ uint(y)
	var onesCount uint8 = 0
	for t > 0 {
		onesCount += uint8(t & 1)
		t = t >> 1
	}
	return int(onesCount)
}

// import (
// 	"math/bits"
// )

// func hammingDistance(x int, y int) int {
//     return bits.OnesCount(uint(x) ^ uint(y))
// }
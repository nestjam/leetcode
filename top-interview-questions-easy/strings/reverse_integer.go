package strings

import (
	"strconv"
)

func reverse(x int) int {
	isNegative := x < 0
	if (isNegative){
		x = -x
	}

	s := strconv.Itoa(x)
	t := []byte(s)
	reverseSlice(t)
	i, err := strconv.ParseInt(string(t), 10, 32)
	
	if (err != nil){
		i = 0
	}
	if (isNegative){
		i = -i
	}
	return int(i)
}

func reverseSlice(s []byte)  {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
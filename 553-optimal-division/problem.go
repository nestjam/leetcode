package optimaldivision

import (
	"strconv"
	"strings"
)

func optimalDivision(nums []int) string {
	b := strings.Builder{}

	b.WriteString(strconv.Itoa(nums[0]))

	if len(nums) < 3 {
		for i := 1; i < len(nums); i++ {
			b.WriteString("/")
			b.WriteString(strconv.Itoa(nums[i]))
		}
	} else {
		b.WriteString("/(")

		b.WriteString(strconv.Itoa(nums[1]))

		for i := 2; i < len(nums); i++ {
			b.WriteString("/")
			b.WriteString(strconv.Itoa(nums[i]))
		}

		b.WriteString(")")
	}

	return b.String()
}

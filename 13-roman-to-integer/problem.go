package romantointeger

import "strings"

func romanToInt(s string) int {
	roman := [...]string{
		"M",
		"CM",
		"D",
		"CD",
		"C",
		"XC",
		"L",
		"XL",
		"X",
		"IX",
		"V",
		"IV",
		"I",
	}
	arabic := [...]int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}

	value := 0
	for i := 0; i < len(roman); i++ {
		for strings.HasPrefix(s, roman[i]) {
			value += arabic[i]
			s = strings.TrimPrefix(s, roman[i])
		}
	}
	return value
}

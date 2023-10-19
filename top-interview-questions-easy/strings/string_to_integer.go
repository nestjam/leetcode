package strings

func myAtoi(s string) int {
	const code0 = 48
	const code9 = 57
	negativeSign := false
	signRead := false
	var num []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && num == nil && !signRead {
			continue
		} else if s[i] == '+' && !signRead && num == nil {
			signRead = true
			continue
		} else if s[i] == '-' && !signRead && num == nil {
			signRead = true
			negativeSign = true
		} else if s[i] >= code0 && s[i] <= code9 {
			num = append(num, s[i])
		} else {
			break
		}
	}
	
	var number int
	const minInt32 uint64 = 2147483648
	var n uint64
	s = omitLeadingZeroes(string(num));
	if len(s) > 10 {
		n = minInt32
	} else {
		n = parseNumber(s)
	}
	
	if negativeSign {
		if n > minInt32 {
			number = -2147483648
		} else { 
			number = -int(n)
		}
	} else {
		const maxInt32 uint64 = 2147483647
		if n > maxInt32 {
			number = int(maxInt32)
		} else { 
			number = int(n)
		}
	}
	return number
}

func omitLeadingZeroes(s string) string {
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] != '0' {
			break
		} 
		j++
	}
	return s[j:]
}

func parseNumber(s string) uint64 {
	m := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	var number uint64 = 0
	var p int = 1
	for i := 0; i < len(s); i++ {
		number = number * 10 + uint64(m[s[i]])
		p *= 10
	}
	return number
}
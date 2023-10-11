package strings

func longestCommonPrefix(strs []string) string {
	var prefix string;
	
	if len(strs) == 0 {
		return prefix
	}

	prefix = strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = getCommonPrefix(prefix, strs[i])
	}
	
	return prefix
}

func getCommonPrefix(a, b string) string {
	l := len(a)
	if len(a) > len(b) {
		l = len(b)
	} 
	i := 0
	for ; i < l; i++ {
		if a[i] != b[i] {
			break
		}
	}
	return string(a[0:i])
}
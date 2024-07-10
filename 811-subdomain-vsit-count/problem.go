package subdomainvsitcount

import (
	"strconv"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)

	for i := 0; i < len(cpdomains); i++ {
		count, domain := split(cpdomains[i], ' ')
		n, _ := strconv.Atoi(count)

		for {
			m[domain] += n

			if domain = after(domain, '.'); domain == "" {
				break
			}
		}
	}

	cpdomains = make([]string, len(m))
	var i int
	for domain, count := range m {
		cpdomains[i] = strconv.Itoa(count) + " " + domain
		i++
	}

	return cpdomains
}

func split(s string, c byte) (left, right string) {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return s[:i], s[min(i+1, len(s)):]
		}
	}
	return s, ""
}

func after(s string, c byte) string {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return s[min(i+1, len(s)):]
		}
	}
	return ""
}

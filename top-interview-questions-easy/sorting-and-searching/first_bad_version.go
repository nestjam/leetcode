package sortingandsearching

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

type validator interface {
	isBadVersion(version int) bool
}

func firstBadVersion(n int, v validator) int {
	p, r := 1, n
	for p < r {
		q := (p + r) / 2
		if v.isBadVersion(q) {
			r = q
		} else {
			p = q + 1
		}
	}
	return r
}

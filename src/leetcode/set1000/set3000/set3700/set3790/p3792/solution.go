package p3792

import "slices"

func reversePrefix(s string, k int) string {
	buf := []byte(s)
	slices.Reverse(buf[:k])
	return string(buf)
}

package p3760

import "math/bits"

func maxDistinct(s string) int {
	var flag int
	for i := range s {
		flag |= 1 << (s[i] - 'a')
	}

	return bits.OnesCount(uint(flag))
}

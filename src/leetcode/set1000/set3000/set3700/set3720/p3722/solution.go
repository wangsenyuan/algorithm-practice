package p3722

import (
	"slices"
	"strings"
)

func lexSmallest(s string) string {
	buf := []byte(s)
	n := len(buf)

	ans := s

	play := func(k int) {
		copy(buf, s)
		slices.Reverse(buf[:k+1])
		if strings.Compare(string(buf), ans) < 0 {
			ans = string(buf)
		}

		slices.Reverse(buf[:k+1])
		slices.Reverse(buf[k+1:])
		if strings.Compare(string(buf), ans) < 0 {
			ans = string(buf)
		}
	}

	for i := range n {
		play(i)
	}

	return ans
}

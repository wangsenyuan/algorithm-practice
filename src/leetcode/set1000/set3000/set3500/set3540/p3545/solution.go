package p3545

import "sort"

func minDeletion(s string, k int) int {
	freq := make([]int, 26)
	for _, c := range []byte(s) {
		freq[c-'a']++
	}
	sort.Ints(freq)
	var sum int
	for i := 25; i >= 0 && k > 0; i-- {
		if freq[i] == 0 {
			break
		}
		sum += freq[i]
		k--
	}
	if k > 0 {
		return 0
	}
	return len(s) - sum
}

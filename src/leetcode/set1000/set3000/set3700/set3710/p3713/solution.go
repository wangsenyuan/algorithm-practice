package p3713

func longestBalanced(s string) int {
	freq := make([]int, 26)

	var res int
	for i := 0; i < len(s); i++ {
		clear(freq)
		for j := i; j < len(s); j++ {
			x := int(s[j] - 'a')
			freq[x]++
			if check(freq) {
				res = max(res, j-i+1)
			}
		}
	}
	return res
}

func check(freq []int) bool {
	var u int
	for _, v := range freq {
		if v == 0 {
			continue
		}
		if u == 0 {
			u = v
		} else if u != v {
			return false
		}
	}
	return true
}

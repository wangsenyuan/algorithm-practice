package p3803

func residuePrefixes(s string) int {
	n := len(s)
	var res int

	freq := make([]int, 26)
	var cnt int

	for i := range n {
		x := int(s[i] - 'a')
		freq[x]++
		if freq[x] == 1 {
			cnt++
			if cnt > 2 {
				break
			}
		}
		if cnt == (i+1)%3 {
			res++
		}
	}

	return res
}

package p3853

func mergeCharacters(s string, k int) string {
	n := len(s)
	last := make([]int, 26)
	for i := range 26 {
		last[i] = -1
	}
	var res []byte
	for i := range n {
		v := int(s[i] - 'a')
		if last[v] >= 0 && len(res)-last[v] <= k {
			continue
		}
		res = append(res, s[i])
		last[v] = len(res) - 1
	}

	return string(res)
}

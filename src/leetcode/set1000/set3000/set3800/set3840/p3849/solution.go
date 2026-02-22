package p3849

func maximumXor(s string, t string) string {
	freq := []int{0, 0}
	for _, c := range t {
		freq[int(c-'0')]++
	}
	var buf []byte
	for i := 0; i < len(s); i++ {
		c := int(s[i] - '0')
		if freq[c^1] > 0 {
			buf = append(buf, '1')
			freq[c^1]--
		} else {
			buf = append(buf, '0')
			freq[c]--
		}
	}
	return string(buf)
}

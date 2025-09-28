package p3692

func majorityFrequencyGroup(s string) string {
	f1 := make([]int, 26)
	for i := range s {
		x := int(s[i] - 'a')
		f1[x]++
	}

	f2 := make(map[int]int)
	for _, v := range f1 {
		f2[v]++
	}

	x := 1
	for k, v := range f2 {
		if k == 0 {
			continue
		}
		if v > f2[x] || v == f2[x] && k > x {
			x = k
		}
	}
	var buf []byte
	for i, v := range f1 {
		if v == x {
			buf = append(buf, byte(i+'a'))
		}
	}
	return string(buf)
}

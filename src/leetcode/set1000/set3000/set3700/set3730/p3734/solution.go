package p3734

func lexPalindromicPermutation(s string, target string) string {
	n := len(target)
	buf := []byte(target)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		buf[j] = buf[i]
	}
	if string(buf) > target && checkFreq(s, string(buf)) {
		return string(buf)
	}
	freq := make([]int, 26)
	for i := range n {
		x := int(s[i] - 'a')
		freq[x]++
	}
	var odd int
	odd_at := -1
	for i := range 26 {
		if freq[i]%2 == 1 {
			odd++
			odd_at = i
		}
		freq[i] /= 2
	}
	if n%2 == 0 && odd > 0 || n%2 == 1 && odd != 1 {
		// 无法得到回文
		return ""
	}

	// 中间的位置是固定的
	next := make([]int, n/2)
	for i := n/2 - 1; i >= 0; i-- {
		if i == n/2-1 || target[i] != target[i+1] {
			next[i] = i + 1
		} else {
			next[i] = next[i+1]
		}
	}

	f := make([]int, 26)

	check := func(p1 int) bool {
		copy(f, freq)
		v := 25
		for i := p1; i < n/2; {
			for f[v] == 0 {
				v--
			}
			u := int(target[i] - 'a')
			if u < v {
				return true
			}
			if u > v {
				return false
			}
			d := min(f[v], next[i]-i)
			f[v] -= d
			i += d
		}
		if odd_at < 0 {
			return false
		}
		return int(target[n/2]-'a') < odd_at
	}

	if !check(0) {
		// 无法得到一个比target更大的序列
		return ""
	}

	clear(buf)
	// 必须在前半段就产生一个更大的序列
	for i := 0; i < n/2; i++ {
		x := int(target[i] - 'a')
		// 是否可以在这里使用x
		if freq[x] > 0 {
			freq[x]--
			if check(i + 1) {
				buf[i] = byte(x + 'a')
				continue
			}
			freq[x]++
		}
		x++
		for freq[x] == 0 {
			x++
		}
		buf[i] = byte(x + 'a')
		freq[x]--

		x = 0
		for j := i + 1; j < n/2; j++ {
			for freq[x] == 0 {
				x++
			}
			buf[j] = byte(x + 'a')
			freq[x]--
		}
		break
	}

	if odd_at >= 0 {
		buf[n/2] = byte(odd_at + 'a')
	}
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		buf[r] = buf[l]
	}
	return string(buf)
}

func checkFreq(s string, t string) bool {
	freq := make([]int, 26)
	for i := range len(s) {
		freq[int(s[i]-'a')]++
		freq[int(t[i]-'a')]--
	}
	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}

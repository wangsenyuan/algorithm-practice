package p3805

import "slices"

func countPairs(words []string) int64 {
	normalize := func(s string) string {
		buf := []byte(s)
		diff := int(s[0] - 'a')

		for i := range buf {
			x := int(buf[i] - 'a')
			x = (x - diff + 26) % 26
			buf[i] = byte(x + 'a')
		}

		return string(buf)
	}

	n := len(words)
	arr := make([]string, n)

	for i, s := range words {
		arr[i] = normalize(s)
	}

	slices.Sort(arr)

	var res int
	for i := 0; i < n; {
		j := i
		for i < n && arr[i] == arr[j] {
			i++
		}
		res += (i - j) * (i - j - 1) / 2
	}
	return int64(res)
}

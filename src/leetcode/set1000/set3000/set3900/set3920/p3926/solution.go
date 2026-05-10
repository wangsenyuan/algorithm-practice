package p3926

import "bytes"

func countWordOccurrences(chunks []string, queries []string) []int {
	freq := make(map[string]int)

	var buf bytes.Buffer
	var last byte = ' '

	checkAppendable := func(x byte) bool {
		if x >= 'a' && x <= 'z' {
			return true
		}
		if x == '-' && (buf.Len() > 0 && last >= 'a' && last <= 'z') {
			return true
		}
		return false
	}

	add := func() {
		x := buf.String()
		// not appendable
		if last == '-' {
			x = x[:len(x)-1]
		}
		if len(x) > 0 {
			freq[x]++
		}
		buf.Reset()
		last = ' '
	}

	for _, chunk := range chunks {
		for i := 0; i < len(chunk); i++ {
			if checkAppendable(chunk[i]) {
				buf.WriteByte(chunk[i])
				last = chunk[i]
				continue
			}
			if buf.Len() > 0 {
				add()
			}
		}
	}

	if buf.Len() > 0 {
		add()
	}

	res := make([]int, len(queries))
	for i, query := range queries {
		res[i] = freq[query]
	}
	return res
}

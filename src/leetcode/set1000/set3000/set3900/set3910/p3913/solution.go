package p3913

import (
	"cmp"
	"slices"
	"strings"
)

type data struct {
	id  int
	pos int
	val byte
}

func sortVowels(s string) string {
	vowels := "aeiou"
	var buf []data
	freq := make([]int, 5)
	var ww []int
	first := make([]int, 5)
	n := len(s)

	for i := n - 1; i >= 0; i-- {
		j := strings.IndexByte(vowels, s[i])
		if j >= 0 {
			buf = append(buf, data{j, i, s[i]})
			freq[j]++
			ww = append(ww, i)
			first[j] = i
		}
	}

	var ids []int
	for i := range 5 {
		if freq[i] > 0 {
			ids = append(ids, i)
		}
	}

	if len(ids) == 0 {
		return s
	}

	slices.SortFunc(ids, func(a int, b int) int {
		return cmp.Or(freq[b]-freq[a], first[a]-first[b])
	})

	pos := make([]int, 5)
	for i, v := range ids {
		pos[v] = i
	}

	slices.SortFunc(buf, func(a data, b data) int {
		if a.id == b.id {
			return a.pos - b.pos
		}
		return pos[a.id] - pos[b.id]
	})

	slices.Reverse(ww)

	res := []byte(s)
	for i, j := range ww {
		res[j] = buf[i].val
	}

	return string(res)
}

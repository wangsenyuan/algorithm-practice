package p3775

import (
	"bytes"
	"slices"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	countVowels := func(word string) int {
		var res int
		for _, c := range word {
			if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
				res++
			}
		}
		return res
	}

	cnt0 := countVowels(words[0])

	var buf bytes.Buffer
	buf.WriteString(words[0])
	for i := 1; i < len(words); i++ {
		buf.WriteByte(' ')
		cnt1 := countVowels(words[i])
		if cnt1 == cnt0 {
			tmp := []byte(words[i])
			slices.Reverse(tmp)
			buf.WriteString(string(tmp))
		} else {
			buf.WriteString(words[i])
		}
	}

	return buf.String()
}

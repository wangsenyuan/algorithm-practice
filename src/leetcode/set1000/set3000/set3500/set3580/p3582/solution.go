package p3582

import (
	"bytes"
	"strings"
)

func generateTag(caption string) string {
	s := strings.ToLower(caption)
	words := strings.Split(s, " ")
	var buf bytes.Buffer
	buf.WriteString("#")
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		if buf.Len() == 1 {
			buf.WriteString(word)
		} else {
			buf.WriteString(strings.ToUpper(word[:1]))
			buf.WriteString(word[1:])
		}
	}
	if buf.Len() > 100 {
		buf.Truncate(100)
	}
	return buf.String()
}

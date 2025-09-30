package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	fmt.Println(res[0])
	fmt.Println(res[1])
}

func solve(s string) []string {

	buf := make([]byte, len(s)*2+1)
	var m int
	if isDelimiter(s[0]) {
		buf[m] = '>'
		m++
	}

	for i := 0; i < len(s); i++ {
		buf[m] = s[i]
		m++
		if isDelimiter(s[i]) && (i == len(s)-1 || isDelimiter(s[i+1])) {
			buf[m] = '>'
			m++
		}
	}

	s = string(buf[:m])

	var a bytes.Buffer
	var b bytes.Buffer

	for i := 0; i < len(s); i++ {
		// 这个肯定是一个开始，找到下一个,;
		j := i
		ok := true
		for i < len(s) && !isDelimiter(s[i]) {
			if s[i] < '0' || s[i] > '9' {
				ok = false
			}
			i++
		}
		if j < i && ok && (j+1 == i || s[j] != '0') {
			w := s[j:i]
			for len(w) > 0 && w[0] == '0' {
				w = w[1:]
			}
			if len(w) == 0 {
				w = "0"
			}
			if a.Len() > 0 {
				a.WriteByte(',')
			}
			a.WriteString(w)
		} else {
			if b.Len() > 0 {
				b.WriteByte(',')
			}
			b.WriteString(s[j:i])
		}
	}

	rewrite := func(buf bytes.Buffer) string {
		if buf.Len() == 0 {
			return "-"
		}
		s := buf.String()
		s = strings.ReplaceAll(s, ">", "")
		return "\"" + s + "\""
	}

	return []string{rewrite(a), rewrite(b)}
}

func isDelimiter(c byte) bool {
	return c == ',' || c == ';'
}

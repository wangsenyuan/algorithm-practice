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
	s := readString(reader)
	res := solve(s)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) string {
	// 只在需要的时候进行分割
	n := len(s)
	var buf bytes.Buffer
	var prev int
	for i := 1; i+1 < n; i++ {
		if isConstant(s[i-1]) && isConstant(s[i]) && isConstant(s[i+1]) {
			if s[i-1] != s[i] || s[i] != s[i+1] || s[i-1] != s[i+1] {
				buf.WriteString(s[prev : i+1])
				buf.WriteByte(' ')
				prev = i + 1
			}
		}
	}

	if prev < n {
		buf.WriteString(s[prev:])
		buf.WriteByte(' ')
	}
	buf.Truncate(buf.Len() - 1)
	return buf.String()
}

func isConstant(c byte) bool {
	return c != 'a' && c != 'e' && c != 'i' && c != 'o' && c != 'u'
}

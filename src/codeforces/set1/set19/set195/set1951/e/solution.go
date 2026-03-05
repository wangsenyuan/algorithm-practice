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

	tc := readNum(reader)

	var buf bytes.Buffer
	for range tc {
		s := readString(reader)
		res := solve(s)

		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			fmt.Fprintf(&buf, "YES\n%d\n", len(res))
			for _, x := range res {
				buf.WriteString(x)
				buf.WriteByte(' ')
			}
			buf.WriteByte('\n')
		}
	}
	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func solve(s string) []string {
	if !checkPalindrome(s) {
		return []string{s}
	}

	for i := range len(s) {
		if s[i] != s[0] {
			if !checkPalindrome(s[i+1:]) {
				return []string{s[:i+1], s[i+1:]}
			}
			if i == 1 || i == len(s)/2 {
				break
			}
			return []string{s[:i+2], s[i+2:]}
		}
	}

	return nil
}

func checkPalindrome(s string) bool {
	n := len(s)
	for i := range n / 2 {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

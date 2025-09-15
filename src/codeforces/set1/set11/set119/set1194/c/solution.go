package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for range tc {
		res := drive(reader)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) bool {
	s := readString(reader)
	t := readString(reader)
	p := readString(reader)
	return solve(s, t, p)
}

func solve(s string, t string, p string) bool {
	freq := make([]int, 26)
	for i := range len(p) {
		freq[int(p[i]-'a')]++
	}

	var j int
	for i := 0; i < len(t); i++ {
		if j < len(s) && t[i] == s[j] {
			j++
			continue
		}
		x := int(t[i] - 'a')
		if freq[x] == 0 {
			return false
		}
		freq[x]--
	}
	return j == len(s)
}

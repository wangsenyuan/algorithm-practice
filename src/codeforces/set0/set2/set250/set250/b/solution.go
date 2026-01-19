package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	n, _ := strconv.Atoi(s)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for range n {
		s = readString(reader)
		res := solve(s)
		fmt.Fprintln(writer, res)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) string {
	// 总共8段

	play := func(x string) string {
		var buf []byte
		n := len(x)
		for i := 0; i < n; i++ {
			j := i
			for i < n && x[i] != ':' {
				i++
			}
			w := string(x[j:i])
			if i-j < 4 {
				w = strings.Repeat("0", 4-len(w)) + w
			}
			if len(buf) > 0 {
				buf = append(buf, ':')
			}
			buf = append(buf, []byte(w)...)
		}
		return string(buf)
	}

	for i := 0; i+1 < len(s); i++ {
		if s[i] == ':' && s[i+1] == ':' {
			res := play(s[:i])
			var res2 string
			if i+2 < len(s) {
				res2 = play(s[i+2:])
				// 8 * 4 + 7
				for len(res)+len(res2) < 8*4+6 {
					res += ":0000"
				}
				res += ":" + res2
			} else {
				for len(res)+len(res2) < 8*4+7 {
					res += ":0000"
				}
			}
			if len(res) > 8*4+7 {
				res = res[1:]
			}

			return res
		}
	}

	// 没有 ::
	return play(s)
}

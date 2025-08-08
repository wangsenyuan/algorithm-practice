package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := process(reader)
	if len(res) == 0 {
		fmt.Println("No such string")
		return
	}
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) (string, string, string) {
	s := readString(reader)
	t := readString(reader)
	return s, t, solve(s, t)
}

func solve(s string, t string) string {
	buf := []byte(s)
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == t[i] {
			continue
		}
		// s[i] < t[i]
		if int(s[i]-'a')+1 < int(t[i]-'a') {
			buf[i]++
			return string(buf)
		}
		// s[i] + 1 = t[j]
		// 看能否使用 s[i]zzzz, 或者 t[i]aaaaa
		for j := i + 1; j < n; j++ {
			buf[j] = 'z'
		}
		if s < string(buf) {
			return string(buf)
		}
		buf[i] = t[i]
		for j := i + 1; j < n; j++ {
			buf[j] = 'a'
		}
		if string(buf) < t {
			return string(buf)
		}
		break
	}
	return ""
}

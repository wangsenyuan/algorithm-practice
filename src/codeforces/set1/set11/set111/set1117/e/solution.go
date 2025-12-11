package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ask := func(s string) string {
		fmt.Fprintln(os.Stdout, "?", s)
		return readString(reader)
	}

	t := readString(reader)
	res := solve(t, ask)
	fmt.Fprintln(os.Stdout, "!", res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(t string, ask func(string) string) string {
	n := len(t)
	s := make([][]byte, 3)
	for i := range 3 {
		s[i] = make([]byte, n)
	}

	w := make([]string, 3)
	for i := range n {
		s[0][i] = byte('a' + i%26)
		s[1][i] = byte('a' + (i/26)%26)
		s[2][i] = byte('a' + (i/26/26)%26)
	}

	for i := range 3 {
		w[i] = ask(string(s[i]))
	}

	p := make([]int, n)
	for i := range n {
		p[i] = int(w[0][i]-'a') + int(w[1][i]-'a')*26 + int(w[2][i]-'a')*26*26
	}

	res := make([]byte, n)

	for i := range n {
		res[p[i]] = t[i]
	}

	return string(res)
}

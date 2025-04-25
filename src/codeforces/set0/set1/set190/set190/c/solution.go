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
	readString(reader)
	s := readString(reader)
	res := solve(s)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

const fail = "Error occurred"

type state struct {
	s   string
	cnt int
}

func solve(s string) string {
	words := strings.Split(strings.TrimSpace(s), " ")

	m := len(words)

	var buf bytes.Buffer

	var dfs func(i int) int

	dfs = func(i int) int {
		if i >= m {
			return -1
		}
		if words[i] == "pair" {
			buf.WriteString("pair<")
			j := dfs(i + 1)
			if j == m || j < 0 {
				return -1
			}
			buf.WriteByte(',')
			j = dfs(j)
			if j < 0 {
				return -1
			}
			buf.WriteByte('>')
			return j
		}
		buf.WriteString("int")
		return i + 1
	}

	p := dfs(0)

	if p != m {
		return fail
	}

	return buf.String()
}

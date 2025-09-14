package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()
	reader := bufio.NewReader(r)
	_, res := drive(reader)
	if len(res) == 0 {
		fmt.Fprintln(w, -1)
		return
	}
	var buf bytes.Buffer
	for _, s := range res {
		buf.WriteString(s)
		buf.WriteByte('\n')
	}
	buf.WriteTo(w)
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

func drive(reader *bufio.Reader) (words []string, res []string) {
	n := readNum(reader)
	words = make([]string, n)
	for i := range n {
		words[i] = readString(reader)
	}
	res = solve(words)
	return
}

func solve(words []string) []string {
	n := len(words)

	abbr := make(map[string]int)
	var short []string

	buf := make([]byte, 4)
	var dfs func(id int, pos int, buf_pos int)

	adj := make([][]int, n)

	connect := func(u int, s string) {
		if _, ok := abbr[s]; !ok {
			short = append(short, s)
			abbr[s] = len(abbr)
		}
		v := abbr[s]
		adj[u] = append(adj[u], v)
	}

	dfs = func(id int, pos int, buf_pos int) {
		if buf_pos > 0 {
			connect(id, string(buf[:buf_pos]))
		}
		if pos == len(words[id]) {
			return
		}

		dfs(id, pos+1, buf_pos)
		if buf_pos < 4 {
			buf[buf_pos] = words[id][pos]
			dfs(id, pos+1, buf_pos+1)
		}
	}

	for i := range n {
		dfs(i, 0, 0)
	}

	m := len(abbr)
	pair := make([]int, m)
	for i := range m {
		pair[i] = -1
	}
	marked := make([]bool, m)
	use := make([]int, n)

	var dfs2 func(u int) bool

	dfs2 = func(u int) bool {
		for _, v := range adj[u] {
			if !marked[v] {
				marked[v] = true
				if pair[v] == -1 || dfs2(pair[v]) {
					use[u] = v
					pair[v] = u
					return true
				}
			}
		}
		return false
	}

	for i := range n {
		use[i] = -1
		clear(marked)
		if !dfs2(i) {
			return nil
		}
	}

	ans := make([]string, n)
	for i := range n {
		ans[i] = short[use[i]]
	}

	return ans
}

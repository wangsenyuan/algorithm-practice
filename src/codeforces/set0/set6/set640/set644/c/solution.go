package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	fmt.Fprintln(writer, len(res))
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) []string {
	s := readString(reader)
	n, _ := strconv.Atoi(s)
	hosts := make([]string, n)
	for i := range n {
		hosts[i] = readString(reader)
	}
	return solve(hosts)
}

func solve(hosts []string) []string {
	queries := make(map[string][]string)

	add := func(host string, path string) {
		if _, found := queries[host]; !found {
			queries[host] = []string{path}
		} else {
			queries[host] = append(queries[host], path)
		}
	}

	for _, cur := range hosts {
		// 按照 http://<host>[/path] 进行分割
		w := cur[7:]
		var i int
		for i < len(w) && w[i] != '/' {
			i++
		}
		add(w[:i], w[i:])
	}

	tr := NewTrie()

	var arr []string

	for host, qs := range queries {
		slices.Sort(qs)
		qs = slices.Compact(qs)
		var buf []byte
		for _, q := range qs {
			if len(q) > 0 {
				buf = append(buf, q...)
			}
			// 即使是空串, 也需要
			buf = append(buf, '$')
		}
		tr.Add(string(buf), len(arr))
		arr = append(arr, host)
	}

	var res []string

	for i := 1; i < len(tr.val); i++ {
		if len(tr.val[i]) > 1 {
			var buf []byte
			for _, id := range tr.val[i] {
				host := "http://" + arr[id]
				if len(buf) > 0 {
					buf = append(buf, ' ')
				}
				buf = append(buf, host...)
			}
			res = append(res, string(buf))
		}
	}

	return res
}

// a...z . /, $

func getId(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a')
	}
	if c == '.' {
		return 26
	}
	if c == '/' {
		return 27
	}
	return 28
}

type Trie struct {
	children [][29]int
	val      [][]int
}

func NewTrie() *Trie {
	children := make([][29]int, 1)
	val := make([][]int, 1)
	return &Trie{
		children: children,
		val:      val,
	}
}

func (tr *Trie) next() int {
	tr.children = append(tr.children, [29]int{})
	tr.val = append(tr.val, []int{})
	return len(tr.children) - 1
}

func (tr *Trie) Add(s string, id int) {
	var node int
	for i := range s {
		c := getId(s[i])
		if tr.children[node][c] == 0 {
			tr.children[node][c] = tr.next()
		}
		node = tr.children[node][c]
	}
	tr.val[node] = append(tr.val[node], id)
}

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
	res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	words := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &words[i])
	}
	patterns := make([]string, m)
	for i := range m {
		fmt.Fscan(reader, &patterns[i])
	}
	return solve(words, patterns)
}

func solve(words []string, patterns []string) []int {
	tr := NewTrie()

	for _, w := range words {
		tr.AddWord(w)
	}

	marked := make([]bool, len(tr.next))

	var dfs func(s string, i int, u int, add bool) int

	dfs = func(s string, i int, u int, add bool) int {
		if u == 0 {
			return 0
		}
		if i == len(s) {
			if add {
				if !marked[u] {
					marked[u] = true
					return tr.cnt[u]
				}
				return 0
			}
			marked[u] = false
			return 0
		}

		if s[i] == '?' {
			res := dfs(s, i+1, u, add)
			for x := range 5 {
				res += dfs(s, i+1, tr.next[u][x], add)
			}
			return res
		}
		x := int(s[i] - 'a')
		return dfs(s, i+1, tr.next[u][x], add)
	}

	res := make([]int, len(patterns))

	for i, p := range patterns {
		res[i] = dfs(p, 0, 1, true)
		dfs(p, 0, 1, false)
	}

	return res
}

type Trie struct {
	next [][5]int
	cnt  []int
}

func NewTrie() *Trie {
	return &Trie{next: make([][5]int, 2), cnt: make([]int, 2)}
}

func (t *Trie) AddWord(s string) {
	cur := 1
	for i := range len(s) {
		x := int(s[i] - 'a')
		if t.next[cur][x] == 0 {
			t.next[cur][x] = len(t.next)
			t.next = append(t.next, [5]int{})
			t.cnt = append(t.cnt, 0)
		}
		cur = t.next[cur][x]
	}
	t.cnt[cur]++
}

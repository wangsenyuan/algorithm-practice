package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := strings.Join(res, "\n")
	fmt.Println(s)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func drive(reader *bufio.Reader) []string {
	n, m := readTwoNums(reader)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = readString(reader)
	}
	queries := make([]string, m)
	for i := 0; i < m; i++ {
		queries[i] = readString(reader)
	}
	return solve(words, queries)
}

type data struct {
	pos  int
	c    int
	pref int
	suf  int
}

func solve(words []string, queries []string) []string {
	pref := NewTrie()
	suf := NewTrie()

	mem := make(map[data]bool)

	var m int
	for _, word := range words {
		m = max(m, len(word))
	}

	dp := make([]data, m+1)

	for _, s := range words {
		for i := range len(s) {
			dp[i].pref = 0
			dp[i].suf = 0
		}
		var node int
		for i := 0; i < len(s); i++ {
			dp[i] = data{
				pos:  i,
				c:    int(s[i] - 'a'),
				pref: node,
			}
			node = pref.Add(node, int(s[i]-'a'))
		}

		node = 0
		for i := len(s) - 1; i >= 0; i-- {
			dp[i].suf = node
			mem[dp[i]] = true
			node = suf.Add(node, int(s[i]-'a'))
		}
	}

	check := func(s string) bool {
		if len(s) > m {
			return false
		}
		for i := range len(s) {
			dp[i].pref = -1
			dp[i].suf = -1
		}
		var node int

		for i := range len(s) {
			dp[i] = data{
				pos:  i,
				c:    int(s[i] - 'a'),
				pref: node,
			}
			x := int(s[i] - 'a')
			node = pref.nodes[node][x]
			if node == 0 {
				// can't proceed
				break
			}
		}

		node = 0
		for i := len(s) - 1; i >= 0; i-- {
			dp[i].suf = node
			x := int(s[i] - 'a')
			for y := range 3 {
				if x != y {
					dp[i].c = y
					if mem[dp[i]] {
						return true
					}
				}
			}
			node = suf.nodes[node][x]
			if node == 0 {
				break
			}
		}
		return false
	}

	ans := make([]string, len(queries))

	for i, cur := range queries {
		ok := check(cur)
		if ok {
			ans[i] = "YES"
		} else {
			ans[i] = "NO"
		}
	}
	return ans
}

type Trie struct {
	id    int
	nodes [][3]int
}

func NewTrie() *Trie {
	nodes := make([][3]int, 1)
	return &Trie{
		id:    1,
		nodes: nodes,
	}
}

func (tr *Trie) Add(node int, c int) int {
	if tr.nodes[node][c] == 0 {
		tr.nodes[node][c] = tr.id
		if tr.id == len(tr.nodes) {
			tr.nodes = append(tr.nodes, [3]int{0, 0, 0})
			tr.id++
		}
	}
	return tr.nodes[node][c]
}

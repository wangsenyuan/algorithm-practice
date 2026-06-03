package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, l int
	fmt.Fscan(reader, &n, &l)
	s := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &s[i])
	}

	return solve(s, l)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a int, b int) int {
	return a * b % mod
}

func solve(s []string, L int) int {
	var tr AhoCorasick
	tr.expand()
	for i, x := range s {
		tr.Add(x, i)
	}

	tr.Build()
	m := len(tr.node)
	n := len(s)
	T := 1 << n
	dp := make([][]int, T)
	ndp := make([][]int, T)
	for i := range T {
		dp[i] = make([]int, m)
		ndp[i] = make([]int, m)
	}
	dp[0][0] = 1
	for range L {
		for mask := range T {
			for v := range m {
				if dp[mask][v] > 0 {
					for i := range 26 {
						to := tr.node[v][i]
						mask1 := mask | tr.last[to]
						ndp[mask1][to] = add(ndp[mask1][to], dp[mask][v])
					}
				}
			}
		}

		for mask := range T {
			copy(dp[mask], ndp[mask])
			clear(ndp[mask])
		}
	}
	var res int
	for v := range m {
		res = add(res, dp[T-1][v])
	}

	return res
}

type AhoCorasick struct {
	node [][26]int
	last []int
}

func (tr *AhoCorasick) expand() int {
	tr.node = append(tr.node, [26]int{})
	tr.last = append(tr.last, 0)
	return len(tr.node) - 1
}

func (tr *AhoCorasick) Add(s string, id int) {
	var cur int
	for _, c := range s {
		x := int(c - 'a')
		if tr.node[cur][x] == 0 {
			tr.node[cur][x] = tr.expand()
		}
		cur = tr.node[cur][x]
	}
	tr.last[cur] |= 1 << id
}

func (tr *AhoCorasick) Build() {
	link := make([]int, len(tr.node))
	var que []int
	for i := range 26 {
		if tr.node[0][i] != 0 {
			que = append(que, tr.node[0][i])
		}
	}
	for len(que) > 0 {
		v := que[0]
		que = que[1:]
		tr.last[v] |= tr.last[link[v]]
		for i := range 26 {
			u := tr.node[v][i]
			if u == 0 {
				tr.node[v][i] = tr.node[link[v]][i]
			} else {
				link[u] = tr.node[link[v]][i]
				que = append(que, u)
			}
		}
	}
}

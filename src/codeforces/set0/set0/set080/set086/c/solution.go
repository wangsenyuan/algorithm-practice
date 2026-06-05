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
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	s := make([]string, m)
	for i := range m {
		fmt.Fscan(reader, &s[i])
	}
	return solve(n, s)
}

const mod = 1000000009

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(n int, s []string) int {
	var ac AhoCorasick
	ac.expand()

	for _, x := range s {
		if len(x) <= n {
			ac.Add(x)
		}
	}
	ac.Build()

	m := len(ac.node)
	maxLen := 0
	for _, x := range s {
		if len(x) <= n {
			maxLen = max(maxLen, len(x))
		}
	}
	if maxLen == 0 {
		return 0
	}

	dp := make([][]int, m)
	ndp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, maxLen+1)
		ndp[i] = make([]int, maxLen+1)
	}
	dp[0][0] = 1

	for pos := 0; pos < n; pos++ {
		for v := range m {
			for gap := 0; gap <= maxLen; gap++ {
				if dp[v][gap] == 0 {
					continue
				}
				for c := range X {
					to := ac.node[v][c]
					ngap := gap + 1
					if ac.best[to] >= ngap {
						ngap = 0
					}
					if ngap <= maxLen {
						ndp[to][ngap] = add(ndp[to][ngap], dp[v][gap])
					}
				}
			}
		}

		for v := range m {
			copy(dp[v], ndp[v])
			clear(ndp[v])
		}
	}

	var res int
	for v := range m {
		res = add(res, dp[v][0])
	}
	return res
}

const X = 4
const DNA = "ACGT"

type AhoCorasick struct {
	node [][X]int
	link []int
	best []int
}

func (tr *AhoCorasick) expand() int {
	tr.node = append(tr.node, [X]int{})
	tr.link = append(tr.link, 0)
	tr.best = append(tr.best, 0)
	return len(tr.node) - 1
}

func (tr *AhoCorasick) Add(s string) {
	var cur int
	for _, c := range s {
		x := strings.IndexRune(DNA, c)
		if tr.node[cur][x] == 0 {
			tr.node[cur][x] = tr.expand()
		}
		cur = tr.node[cur][x]
	}
	tr.best[cur] = max(tr.best[cur], len(s))
}

func (tr *AhoCorasick) Build() {
	var que []int
	for i := range X {
		if tr.node[0][i] != 0 {
			que = append(que, tr.node[0][i])
		}
	}
	for len(que) > 0 {
		v := que[0]
		que = que[1:]
		tr.best[v] = max(tr.best[v], tr.best[tr.link[v]])
		for i := range X {
			u := tr.node[v][i]
			if u == 0 {
				tr.node[v][i] = tr.node[tr.link[v]][i]
			} else {
				tr.link[u] = tr.node[tr.link[v]][i]
				que = append(que, u)
			}
		}
	}
}

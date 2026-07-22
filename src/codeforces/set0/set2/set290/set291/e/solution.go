package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int64 {
	var n int
	fmt.Fscan(reader, &n)
	parents := make([]int, n-1)
	edges := make([]string, n-1)
	for v := range n - 1 {
		fmt.Fscan(reader, &parents[v], &edges[v])
	}
	var t string
	fmt.Fscan(reader, &t)
	return solve(n, parents, edges, t)
}

func solve(n int, parents []int, edges []string, t string) int64 {
	adj := make([][]int, n)
	for v := 0; v < n-1; v++ {
		adj[parents[v]-1] = append(adj[parents[v]-1], v+1)
	}

	transition := kmpAutomaton(t)

	var res int64

	var dfs2 func(u int, state int)
	dfs2 = func(u int, state int) {
		for _, v := range adj[u] {
			cur := state
			s := edges[v-1]
			for i := 0; i < len(s); i++ {
				code := transition[cur][s[i]-'a']
				if code < 0 {
					res++
					cur = int(-code - 1)
				} else {
					cur = int(code)
				}
			}
			dfs2(v, cur)
		}
	}

	dfs2(0, 0)

	return res
}

// kmpAutomaton precomputes every KMP state transition. A negative value
// encodes a completed match and the state to continue from as -state-1.
func kmpAutomaton(s string) [][26]int32 {
	pt := kmp(s)
	transition := make([][26]int32, len(s))

	for state := 0; state < len(s); state++ {
		for c := 0; c < 26; c++ {
			if s[state] != byte('a'+c) {
				if state > 0 {
					transition[state][c] = transition[pt[state-1]][c]
				}
				continue
			}

			next := state + 1
			if next == len(s) {
				transition[state][c] = int32(-pt[state] - 1)
			} else {
				transition[state][c] = int32(next)
			}
		}
	}

	return transition
}

func kmp(s string) []int {
	n := len(s)
	p := make([]int, n)
	for i := 1; i < n; i++ {
		j := p[i-1]
		for j > 0 && s[i] != s[j] {
			j = p[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		p[i] = j
	}
	return p
}

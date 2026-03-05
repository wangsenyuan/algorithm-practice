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
	res := drive(reader)
	fmt.Println(res[0], res[1])
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) []int {
	readString(reader)
	s := readString(reader)
	_n := readString(reader)
	n, _ := strconv.Atoi(_n)
	synonyms := make([]string, n)
	for i := range n {
		synonyms[i] = readString(reader)
	}
	return solve(s, synonyms)
}

func solve(s string, synonyms []string) []int {
	pos := make(map[string]int)
	var cntR []int
	var wordLen []int

	getOrAdd := func(w string) int {
		w = strings.ToLower(w)
		if v, ok := pos[w]; ok {
			return v
		}
		pos[w] = len(pos)
		cntR = append(cntR, strings.Count(w, "r"))
		wordLen = append(wordLen, len(w))
		return pos[w]
	}

	for _, cur := range synonyms {
		ww := strings.Split(cur, " ")
		getOrAdd(ww[0])
		getOrAdd(ww[1])
	}
	adj := make([][]int, len(wordLen))
	for _, cur := range synonyms {
		ww := strings.Split(cur, " ")
		x := getOrAdd(ww[0])
		y := getOrAdd(ww[1])
		adj[x] = append(adj[x], y)
	}

	n := len(wordLen)
	low := make([]int, n)
	dfn := make([]int, n)
	vis := make([]bool, n)
	stack := make([]int, n)
	var top int
	var timer int

	var comps [][]int
	belong := make([]int, n)

	var dfs func(u int)
	dfs = func(u int) {
		timer++
		dfn[u] = timer
		low[u] = timer
		vis[u] = true
		stack[top] = u
		top++
		for _, v := range adj[u] {
			if dfn[v] == 0 {
				dfs(v)
				low[u] = min(low[u], low[v])
			} else if vis[v] {
				low[u] = min(low[u], dfn[v])
			}
		}
		if low[u] == dfn[u] {
			var cur []int
			for top > 0 {
				v := stack[top-1]
				top--
				vis[v] = false
				cur = append(cur, v)
				belong[v] = len(comps)
				if u == v {
					break
				}
			}
			comps = append(comps, cur)
		}
	}

	for i := range n {
		if dfn[i] == 0 {
			dfs(i)
		}
	}
	dp := make([][]int, len(comps))
	for i, cur := range comps {
		dp[i] = []int{inf, inf}
		for _, v := range cur {
			if cntR[v] < dp[i][0] || cntR[v] == dp[i][0] && wordLen[v] < dp[i][1] {
				dp[i][0] = cntR[v]
				dp[i][1] = wordLen[v]
			}
		}
	}

	tr := make([][]int, len(comps))
	for u := range n {
		for _, v := range adj[u] {
			if belong[u] != belong[v] {
				tr[belong[u]] = append(tr[belong[u]], belong[v])
			}
		}
	}
	que := make([]int, len(comps))
	rev := make([][]int, len(comps))

	var head, tail int

	deg := make([]int, len(comps))
	for i := range len(comps) {
		slices.Sort(tr[i])
		tr[i] = slices.Compact(tr[i])
		for _, v := range tr[i] {
			rev[v] = append(rev[v], i)
		}
		deg[i] = len(tr[i])
		if deg[i] == 0 {
			que[head] = i
			head++
		}
	}

	update := func(u int, v int) {
		if dp[v][0] < dp[u][0] || dp[v][0] == dp[u][0] && dp[v][1] < dp[u][1] {
			dp[u][0] = dp[v][0]
			dp[u][1] = dp[v][1]
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		for _, v := range rev[u] {
			update(v, u)
			deg[v]--
			if deg[v] == 0 {
				que[head] = v
				head++
			}
		}
	}

	res := make([]int, 2)

	ss := strings.Split(s, " ")

	for _, w := range ss {
		w = strings.ToLower(w)
		if v, ok := pos[w]; ok {
			v = belong[v]
			res[0] += dp[v][0]
			res[1] += dp[v][1]
		} else {
			res[0] += strings.Count(w, "r")
			res[1] += len(w)
		}
	}

	return res
}

const inf = 1 << 60

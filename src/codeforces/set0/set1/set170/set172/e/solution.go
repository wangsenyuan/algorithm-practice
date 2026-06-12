package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readInt(reader *bufio.Reader) int {
	s := readString(reader)
	i, _ := strconv.Atoi(s)
	return i
}

func drive(reader *bufio.Reader) []int {
	html := readString(reader)
	n := readInt(reader)

	css := make([]string, n)
	for i := range n {
		css[i] = readString(reader)
	}

	return solve(html, css)
}

func solve(html string, css []string) []int {
	// parse it first
	var nodes []string
	var adj [][]int
	var dp []int

	var freq map[string]int
	var pos int
	n := len(html)
	var dfs func() int
	dfs = func() int {
		if pos >= n || html[pos] != '<' {
			panic(fmt.Errorf("wrong expression process at %d", pos))
		}
		// must be a opentag <a>...</a> or <a/>
		pos++
		i := pos
		for pos < n && (html[pos] >= 'a' && html[pos] <= 'z') {
			pos++
		}
		name := html[i:pos]
		id := len(nodes)
		nodes = append(nodes, name)
		dp = append(dp, 0)
		adj = append(adj, nil)

		if html[pos] == '/' {
			// <a/>
			dp[id] = 1
			freq[name]++
			// 消耗掉 >
			pos += 2
		} else {
			// <a> ... </a>
			keep := freq[name]
			freq[name]++
			pos++
			end := fmt.Sprintf("</%s>", name)

			for html[pos:pos+len(end)] != end {
				adj[id] = append(adj[id], dfs())
			}

			pos += len(end)
			dp[id] = freq[name] - keep
		}

		return id
	}

	var roots []int
	for pos < n {
		freq = make(map[string]int)
		id := dfs()
		roots = append(roots, id)
	}

	m := len(css)
	rules := make([][]string, m)

	for i, cur := range css {
		ss := strings.Split(cur, " ")
		rules[i] = ss
	}

	ans := make([]int, m)

	var play func(u int, rid int, where int)
	play = func(u int, rid int, where int) {
		if nodes[u] == rules[rid][where] {
			where++
		}
		if where == len(rules[rid]) {
			ans[rid] += dp[u]
			return
		}
		for _, v := range adj[u] {
			play(v, rid, where)
		}
	}

	for _, r := range roots {
		for i := range m {
			play(r, i, 0)
		}
	}

	return ans
}

type pair struct {
	first  int
	second int
}

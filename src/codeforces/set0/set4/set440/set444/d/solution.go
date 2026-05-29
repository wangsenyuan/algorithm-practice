package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func drive(reader *bufio.Reader) []int {
	var s string
	fmt.Fscan(reader, &s)
	var n int
	fmt.Fscan(reader, &n)
	queries := make([][]string, n)
	for i := range n {
		queries[i] = make([]string, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(s, queries)
}

func solve(s string, queries [][]string) []int {
	id := make(map[string]int)
	var words []string
	var pos [][]int

	getId := func(x string) int {
		if v, ok := id[x]; ok {
			return v
		}
		v := len(words)
		id[x] = v
		words = append(words, x)
		pos = append(pos, nil)
		return v
	}

	for i := 0; i < len(s); i++ {
		for d := range 4 {
			if i+d == len(s) {
				break
			}
			x := s[i : i+d+1]
			j := getId(x)
			pos[j] = append(pos[j], i)
		}
	}

	n := len(s)
	const inf = 1 << 30
	const threshold = 320

	unionLen := func(i int, a string, j int, b string) int {
		l := min(i, j)
		r := max(i+len(a), j+len(b))
		return r - l
	}

	calcLight := func(a string, b string, ia int, ib int) int {
		u := pos[ia]
		v := pos[ib]
		if len(u) > len(v) {
			u, v = v, u
			a, b = b, a
		}
		best := inf
		for _, i := range u {
			j := sort.SearchInts(v, i)
			if j < len(v) {
				best = min(best, unionLen(i, a, v[j], b))
			}
			if j > 0 {
				best = min(best, unionLen(i, a, v[j-1], b))
			}
		}
		return best
	}

	type ask struct {
		a  int
		b  int
		ok bool
	}

	asks := make([]ask, len(queries))
	heavyNeed := make(map[int]bool)

	for i, cur := range queries {
		ia, oka := id[cur[0]]
		ib, okb := id[cur[1]]
		if !oka || !okb {
			continue
		}
		asks[i] = ask{ia, ib, true}
		if len(pos[ia]) > threshold {
			heavyNeed[ia] = true
		}
		if len(pos[ib]) > threshold {
			heavyNeed[ib] = true
		}
	}

	heavyAns := make(map[int][]int)

	for h := range heavyNeed {
		ans := make([]int, len(words))
		for i := range ans {
			ans[i] = inf
		}

		prev := make([]int, n)
		next := make([]int, n)
		for i := range n {
			prev[i] = -1
			next[i] = n
		}
		for _, p := range pos[h] {
			prev[p] = p
			next[p] = p
		}
		for i := 1; i < n; i++ {
			if prev[i] < 0 {
				prev[i] = prev[i-1]
			}
		}
		for i := n - 2; i >= 0; i-- {
			if next[i] == n {
				next[i] = next[i+1]
			}
		}

		a := words[h]
		for j, b := range words {
			for _, p := range pos[j] {
				if prev[p] >= 0 {
					ans[j] = min(ans[j], unionLen(prev[p], a, p, b))
				}
				if next[p] < n {
					ans[j] = min(ans[j], unionLen(next[p], a, p, b))
				}
			}
		}
		heavyAns[h] = ans
	}

	ans := make([]int, len(queries))
	for i, cur := range asks {
		if !cur.ok {
			ans[i] = -1
			continue
		}
		var best int
		if row, ok := heavyAns[cur.a]; ok {
			best = row[cur.b]
		} else if row, ok := heavyAns[cur.b]; ok {
			best = row[cur.a]
		} else {
			best = calcLight(words[cur.a], words[cur.b], cur.a, cur.b)
		}
		if best == inf {
			best = -1
		}
		ans[i] = best
	}

	return ans
}

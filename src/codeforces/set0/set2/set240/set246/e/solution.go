package main

import (
	"bufio"
	"fmt"
	"os"
)

type Query struct {
	v int
	k int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, x := range drive(reader) {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)

	names := make([]string, n)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &names[i], &parent[i])
	}

	var m int
	fmt.Fscan(reader, &m)
	queries := make([]Query, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &queries[i].v, &queries[i].k)
		queries[i].v--
	}

	return solve(names, parent, queries)
}

func solve(names []string, parent []int, queries []Query) []int {
	n := len(names)
	children := make([][]int, n)
	var roots []int
	for i, p := range parent {
		if p == 0 {
			roots = append(roots, i)
			continue
		}
		children[p-1] = append(children[p-1], i)
	}

	var depthSize []int
	var countDepth func(u int, depth int)
	countDepth = func(u int, depth int) {
		if depth == len(depthSize) {
			depthSize = append(depthSize, 0)
		}
		depthSize[depth]++
		for _, v := range children[u] {
			countDepth(v, depth+1)
		}
	}

	for _, root := range roots {
		countDepth(root, 0)
	}

	bits := make([]BIT, len(depthSize))
	lastSeen := make([]map[string]int, len(depthSize))
	for d, sz := range depthSize {
		bits[d] = NewBIT(sz)
		lastSeen[d] = make(map[string]int)
	}

	queriesAt := make([][]int, n)
	for i, q := range queries {
		queriesAt[q.v] = append(queriesAt[q.v], i)
	}

	nextPos := make([]int, len(depthSize))
	ans := make([]int, len(queries))

	var dfs func(u int, depth int)
	dfs = func(u int, depth int) {
		start := make([]int, len(queriesAt[u]))
		for i, qi := range queriesAt[u] {
			targetDepth := depth + queries[qi].k
			if targetDepth >= len(depthSize) {
				start[i] = -1
				continue
			}
			start[i] = nextPos[targetDepth]
		}

		pos := nextPos[depth]
		if prev, ok := lastSeen[depth][names[u]]; ok {
			bits[depth].Add(prev, -1)
		}
		lastSeen[depth][names[u]] = pos
		bits[depth].Add(pos, 1)
		nextPos[depth]++

		for _, v := range children[u] {
			dfs(v, depth+1)
		}

		for i, qi := range queriesAt[u] {
			if start[i] < 0 {
				continue
			}
			targetDepth := depth + queries[qi].k
			ans[qi] = bits[targetDepth].RangeSum(start[i], nextPos[targetDepth]-1)
		}
	}

	for _, root := range roots {
		dfs(root, 0)
	}

	return ans
}

type BIT []int

func NewBIT(n int) BIT {
	return make(BIT, n+1)
}

func (bit BIT) Add(pos int, delta int) {
	for pos++; pos < len(bit); pos += pos & -pos {
		bit[pos] += delta
	}
}

func (bit BIT) PrefixSum(pos int) int {
	var res int
	for pos++; pos > 0; pos -= pos & -pos {
		res += bit[pos]
	}
	return res
}

func (bit BIT) RangeSum(left int, right int) int {
	if right < left {
		return 0
	}
	return bit.PrefixSum(right) - bit.PrefixSum(left-1)
}

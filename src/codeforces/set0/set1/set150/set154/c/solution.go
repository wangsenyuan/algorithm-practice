package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int64 {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	friendships := make([][2]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &friendships[i][0], &friendships[i][1])
	}
	return solve(n, friendships)
}

func solve(n int, friendships [][2]int) int64 {
	deg := make([]int, n)
	for _, cur := range friendships {
		u, v := cur[0]-1, cur[1]-1
		deg[u]++
		deg[v]++
	}

	start := make([]int, n+1)
	for i := 0; i < n; i++ {
		start[i+1] = start[i] + deg[i]
	}
	next := append([]int(nil), start...)
	adj := make([]int, 2*len(friendships))

	for _, cur := range friendships {
		u, v := cur[0]-1, cur[1]-1
		adj[next[u]] = v
		next[u]++
		adj[next[v]] = u
		next[v]++
	}

	for i := 0; i < n; i++ {
		slices.Sort(adj[start[i]:start[i+1]])
	}

	keys := make([]signature, n)
	for i := 0; i < n; i++ {
		keys[i] = hashSlice(adj[start[i]:start[i+1]])
	}
	ans := countEqual(keys)

	for i := 0; i < n; i++ {
		keys[i] = hashClosed(i, adj[start[i]:start[i+1]])
	}
	ans += countEqual(keys)

	return ans
}

type signature struct {
	a uint64
	b uint64
}

const (
	base1 uint64 = 1000003
	base2 uint64 = 1000033
	seed1 uint64 = 1469598103934665603
	seed2 uint64 = 1099511628211
)

func addValue(sig signature, v int) signature {
	x := uint64(v + 1)
	sig.a = sig.a*base1 + x
	sig.b = sig.b*base2 + x
	return sig
}

func hashSlice(arr []int) signature {
	sig := signature{seed1, seed2}
	for _, v := range arr {
		sig = addValue(sig, v)
	}
	sig = addValue(sig, len(arr)+1_000_000)
	return sig
}

func hashClosed(u int, arr []int) signature {
	sig := signature{seed1, seed2}
	inserted := false
	for _, v := range arr {
		if !inserted && u < v {
			sig = addValue(sig, u)
			inserted = true
		}
		sig = addValue(sig, v)
	}
	if !inserted {
		sig = addValue(sig, u)
	}
	sig = addValue(sig, len(arr)+1+1_000_000)
	return sig
}

func countEqual(keys []signature) int64 {
	slices.SortFunc(keys, func(a, b signature) int {
		return cmp.Or(int(a.a-b.a), int(a.b-b.b))
	})

	var res int64
	for i := 0; i < len(keys); {
		j := i + 1
		for j < len(keys) && keys[j] == keys[i] {
			j++
		}
		cnt := int64(j - i)
		res += cnt * (cnt - 1) / 2
		i = j
	}
	return res
}

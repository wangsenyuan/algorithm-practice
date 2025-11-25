package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &queries[i])
	}
	return solve(a, queries)
}

func solve(a []int, queries []int) []int {
	slices.Sort(a)
	a = slices.Compact(a)
	// n := len(a)

	tr := new(Trie)
	tr.next()

	for _, v := range a {
		tr.Add(v)
	}

	var ans []int
	var x int
	for _, v := range queries {
		x ^= v
		var node int
		var tmp int
		for i := H - 1; i >= 0; i-- {
			bit := (x >> i) & 1
			if tr.cnt[tr.children[node][bit]] == 1<<i {
				bit ^= 1
				tmp |= 1 << i
			}

			node = tr.children[node][bit]
			if node == 0 {
				break
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}

type Trie struct {
	children [][2]int
	cnt      []int
}

func (tr *Trie) next() int {
	tr.children = append(tr.children, [2]int{})
	tr.cnt = append(tr.cnt, 0)
	return len(tr.children) - 1
}

const H = 30

func (tr *Trie) Add(x int) {
	var node int
	for i := H - 1; i >= 0; i-- {
		bit := (x >> i) & 1
		if tr.children[node][bit] == 0 {
			tr.children[node][bit] = tr.next()
		}
		node = tr.children[node][bit]
		tr.cnt[node]++
	}
}

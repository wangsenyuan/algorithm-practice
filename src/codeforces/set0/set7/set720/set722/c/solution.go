package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := readNNums(reader, n)
	p := readNNums(reader, n)
	return solve(a, p)
}

func solve(a []int, p []int) []int {
	n := len(a)

	ans := make([]int, n)
	var best int

	set := NewDSU(n)
	val := make([]int, n)

	pos := make([]int, n)
	for i := range n {
		p[i]--
		pos[p[i]] = i
	}

	for i := n - 1; i >= 0; i-- {
		ans[i] = best
		id := p[i]
		cur := a[id]
		if id > 0 && pos[id-1] > i {
			l := set.Find(id - 1)
			cur += val[l]
			set.Union(id, id-1)
		}

		if id < n-1 && pos[id+1] > i {
			r := set.Find(id + 1)
			cur += val[r]
			set.Union(id, id+1)
		}

		best = max(best, cur)
		val[set.Find(id)] = cur
	}
	// ans[0] = best

	return ans
}

type DSU struct {
	arr  []int
	cnt  []int
	size int
}

func NewDSU(n int) *DSU {
	set := new(DSU)
	set.arr = make([]int, n)
	set.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	set.size = n
	return set
}

func (set *DSU) Find(u int) int {
	if set.arr[u] != u {
		set.arr[u] = set.Find(set.arr[u])
	}
	return set.arr[u]
}

func (set *DSU) Union(a, b int) bool {
	a = set.Find(a)
	b = set.Find(b)
	if a == b {
		return false
	}
	if set.cnt[a] < set.cnt[b] {
		a, b = b, a
	}
	set.cnt[a] += set.cnt[b]
	set.arr[b] = a
	set.size--
	return true
}

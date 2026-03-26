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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	set := NewDSU(n)
	deg := make([]int, n)
	marked := make([]bool, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		marked[u] = true
		marked[v] = true
		if u != v {
			set.Union(u, v)
			deg[u]++
			deg[v]++
		}
	}

	oddDegCnt := make([]int, n)
	sz := set.size
	var roots []int
	for i := range n {
		if i > 0 && !marked[i] {
			sz--
			continue
		}
		if deg[i]&1 == 1 {
			j := set.Find(i)
			oddDegCnt[j]++
		}
		if set.Find(i) == i {
			roots = append(roots, i)
		}
	}

	var res int
	if sz > 1 {
		// 需要x条边，把所有的部分联通起来
		res += sz
		for _, i := range roots {
			if oddDegCnt[i] > 1 {
				// 需要添加两条边，消耗掉两个奇数deg的点
				oddDegCnt[i] -= 2
			}
			// else，添加2条边，不会改变数量
		}
	}
	// 现在已经连起来的，要么大家都是偶数，所以看有多少个奇数deg的点
	var sumOdd int
	for _, i := range roots {
		sumOdd += oddDegCnt[i]
	}

	return res + sumOdd/2
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

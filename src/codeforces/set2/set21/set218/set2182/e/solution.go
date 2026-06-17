package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	a := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &a[i])
	}
	friends := make([][]int, n)
	for i := range n {
		var x, y, z int
		fmt.Fscan(reader, &x, &y, &z)
		friends[i] = []int{x, y, z}
	}
	return solve(k, a, friends)
}

func solve(k int, a []int, friends [][]int) int {

	slices.Sort(a)
	// 按照x升序排列
	slices.SortFunc(friends, func(f []int, s []int) int {
		return f[0] - s[0]
	})

	var tot int
	n := len(friends)
	var pq IntHeap
	var pos int
	var res int
	for _, w := range a {
		for pos < n && friends[pos][0] <= w {
			tot += friends[pos][2]
			d := friends[pos][2] - friends[pos][1]
			heap.Push(&pq, d)
			pos++
		}
		if len(pq) > 0 {
			// 后续的也没法使用w, 这里优先使用最能省钱的那个
			tot -= heap.Pop(&pq).(int)
			res++
		}
	}

	for pos < n {
		tot += friends[pos][2]
		d := friends[pos][2] - friends[pos][1]
		heap.Push(&pq, d)
		pos++
	}

	for tot > k && len(pq) > 0 {
		// 只能让这个朋友不高心,购买y[i]的礼物,且不使用盒子
		tot -= heap.Pop(&pq).(int)
	}

	return res + len(pq)
}

type IntHeap []int

func (pq IntHeap) Len() int {
	return len(pq)
}

func (pq IntHeap) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq IntHeap) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *IntHeap) Push(x any) {
	*pq = append(*pq, x.(int))
}

func (pq *IntHeap) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}

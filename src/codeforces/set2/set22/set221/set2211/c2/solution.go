package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(k, a, b)
}

func solve(k int, a []int, b []int) string {
	n := len(a)
	uf := NewUFSet(n)
	for i := k; i < n; i++ {
		if a[i] == a[i-k] {
			uf.Union(i, i-k)
		}
	}

	value := make([]int, n)
	assign := func(pos int, v int) bool {
		root := uf.Find(pos)
		if value[root] != 0 && value[root] != v {
			return false
		}
		value[root] = v
		return true
	}

	for i, v := range b {
		if v > 0 && !assign(i, v) {
			return "NO"
		}
	}

	for i := k; i < n; i++ {
		if a[i] != a[i-k] {
			if !assign(i-k, a[i-k]) || !assign(i, a[i]) {
				return "NO"
			}
		}
	}

	need := make([]int, n+1)
	for i := range k {
		need[a[i]]++
	}
	for i := range k {
		root := uf.Find(i)
		if value[root] > 0 {
			need[value[root]]--
			if need[value[root]] < 0 {
				return "NO"
			}
		}
	}

	return "YES"
}

type UFSet struct {
	arr []int
}

func NewUFSet(n int) *UFSet {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i
	}
	return &UFSet{arr}
}

func (uf *UFSet) Find(x int) int {
	if uf.arr[x] != x {
		uf.arr[x] = uf.Find(uf.arr[x])
	}
	return uf.arr[x]
}

func (uf *UFSet) Union(a int, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}
	uf.arr[pa] = pb
	return true
}

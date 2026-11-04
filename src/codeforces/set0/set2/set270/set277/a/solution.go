package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	langs := make([][]int, n)
	for i := range n {
		var k int
		fmt.Fscan(reader, &k)
		langs[i] = make([]int, k)
		for j := range k {
			fmt.Fscan(reader, &langs[i][j])
		}
	}
	return solve(m, langs)
}

func solve(m int, langs [][]int) int {
	last := make([]int, m)
	for i := range m {
		last[i] = -1
	}
	n := len(langs)
	s := build(n)
	var todo []int
	for id, cur := range langs {
		for _, i := range cur {
			i--
			if last[i] != -1 {
				s.union(last[i], id)
			}
			last[i] = id
		}
		if len(cur) == 0 {
			todo = append(todo, id)
		}
	}

	if len(todo) == n {
		return n
	}

	var ans int
	for i := range n {
		if s.find(i) == i {
			ans++
		}
	}
	// todo 每个人都必须学一门
	return ans - 1
}

type set struct {
	arr []int
	sz  []int
}

func build(n int) *set {
	arr := make([]int, n)
	sz := make([]int, n)
	for i := range n {
		arr[i] = i
		sz[i] = 1
	}
	return &set{arr, sz}
}

func (s *set) find(u int) int {
	v := s.arr[u]
	for v != s.arr[v] {
		v = s.arr[v]
	}
	for u != v {
		u, s.arr[u] = s.arr[u], v
	}
	return v
}

func (s *set) union(u int, v int) bool {
	u = s.find(u)
	v = s.find(v)
	if u == v {
		return false
	}
	if s.sz[u] < s.sz[v] {
		u, v = v, u
	}
	s.sz[u] += s.sz[v]
	s.arr[v] = u
	return true
}

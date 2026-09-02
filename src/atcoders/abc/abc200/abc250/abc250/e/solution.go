package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	for _, ans := range res {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []string {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(a, b, queries)
}

func solve(a []int, b []int, queries [][]int) []string {
	// TODO: solve by hand first.
	first := NewHashContainer()
	second := NewHashContainer()

	nums := slices.Clone(a)
	nums = append(nums, b...)
	slices.Sort(nums)
	nums = slices.Compact(nums)

	type pair struct {
		f uint64
		s uint64
	}

	for _, v := range nums {
		first.Add(v)
		second.Add(v)
	}

	play := func(arr []int) []pair {
		n := len(arr)
		var cur pair
		res := make([]pair, n)

		vis := make(map[int]bool)

		for i, v := range arr {
			if !vis[v] {
				vis[v] = true
				cur.f ^= first.values[v]
				cur.s ^= second.values[v]
			}

			res[i] = cur
		}
		return res
	}

	u := play(a)
	v := play(b)

	ans := make([]string, len(queries))

	for i, cur := range queries {
		x := cur[0] - 1
		y := cur[1] - 1
		if u[x] == v[y] {
			ans[i] = "Yes"
		} else {
			ans[i] = "No"
		}
	}

	return ans
}

type HashContainer struct {
	values map[int]uint64
}

func NewHashContainer() *HashContainer {
	return &HashContainer{
		values: make(map[int]uint64),
	}
}

func (h *HashContainer) Add(v int) {
	h.values[v] = rand.Uint64()
}
